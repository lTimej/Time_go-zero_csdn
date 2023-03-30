package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/common/globalkey"
	"liujun/Time_go-zero_csdn/csdn/im/model"
	"net"
	"net/http"
	"sync"
	"time"

	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/svc"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"

	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/fatih/set.v0"
)

type Message struct {
	TargetId   string `json:"target_id"`
	Type       int64  `json:"type"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
	UserId     string `json:"user_id"`
}
type Node struct {
	Conn          *websocket.Conn //连接
	Addr          string          //客户端地址
	FirstTime     uint64          //首次连接时间
	HeartbeatTime uint64          //心跳时间
	LoginTime     uint64          //登录时间
	DataQueue     chan []byte     //消息
	GroupSets     set.Interface   //好友 / 群
}

// 映射关系
var clientMap map[string]*Node = make(map[string]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

type UserChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserChatLogic {
	return &UserChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserChatLogic) UserChat(w http.ResponseWriter, r *http.Request) {
	// todo: add your logic here and delete this line

	fmt.Println("init goroutine ")
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	fmt.Println(user_id, "~~~~~~~~~~~~~")
	// isvalida := true //checkToke()  待.........
	conn, err := (&websocket.Upgrader{
		//token 校验
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 5 * time.Second,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		return
	}
	//2.获取conn
	currentTime := uint64(time.Now().Unix())
	node := &Node{
		Conn:          conn,
		Addr:          conn.RemoteAddr().String(), //客户端地址
		HeartbeatTime: currentTime,                //心跳时间
		LoginTime:     currentTime,                //登录时间
		DataQueue:     make(chan []byte, 50),
		GroupSets:     set.New(set.ThreadSafe),
	}
	//3. 用户关系
	//4. userid 跟 node绑定 并加锁
	rwLocker.Lock()
	clientMap[user_id] = node
	rwLocker.Unlock()
	//5.完成发送逻辑
	go l.sendProc(node)
	//6.完成接受逻辑
	go l.recvProc(node, user_id)

	go l.udpSendProc()
	go l.udpRecvProc(user_id)
	//7.加入在线用户到缓存
	l.SetUserOnlineInfo("online_"+user_id, []byte(node.Addr), time.Duration(4*time.Hour))
}

func (l *UserChatLogic) SetUserOnlineInfo(key string, val []byte, timeTTL time.Duration) {
	ctx := context.Background()
	l.svcCtx.RedisIm.Set(ctx, key, val, timeTTL)
}

func (l *UserChatLogic) sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			fmt.Println("[ws]sendProc >>>> msg :", string(data))
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func (l *UserChatLogic) recvProc(node *Node, user_id string) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := Message{UserId: user_id}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			fmt.Println(err)
		}
		//心跳检测 msg.Media == -1 || msg.Type == 3
		if msg.Type == 3 {
			currentTime := uint64(time.Now().Unix())
			node.Heartbeat(currentTime)
		} else {
			l.dispatch(data, user_id)
			l.broadMsg(data) //todo 将消息广播到局域网
			fmt.Println("[ws] recvProc <<<<< ", string(data))
		}

	}
}

// 更新用户心跳
func (node *Node) Heartbeat(currentTime uint64) {
	node.HeartbeatTime = currentTime
}

// 后端调度逻辑处理
func (l *UserChatLogic) dispatch(data []byte, user_id string) {
	fmt.Println(string(data), "===========")
	msg := Message{UserId: user_id}
	msg.CreateTime = time.Now().Unix()
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msg, "===========")
	switch msg.Type {
	case 1: //私信
		fmt.Println("dispatch  data ===:", string(data), "-------------", msg.TargetId, msg.Content)
		l.sendMsg(msg.TargetId, data, user_id)
	case 2: //群发
		l.sendGroupMsg(msg.TargetId, data) //发送的群ID ，消息内容
		// case 4: // 心跳
		// 	node.Heartbeat()
		//case 4:
		//
	}
}

var udpsendChan chan []byte = make(chan []byte, 1024)

func (l *UserChatLogic) broadMsg(data []byte) {
	udpsendChan <- data
}

// 完成udp数据发送协程
func (l *UserChatLogic) udpSendProc() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(172, 20, 16, 20),
		Port: 3001,
	})
	defer con.Close()
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case data := <-udpsendChan:
			fmt.Println("udpSendProc  data :", string(data))
			_, err := con.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

}

// 完成udp数据接收协程
func (l *UserChatLogic) udpRecvProc(user_id string) {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3001,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer con.Close()
	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("udpRecvProc  data :", string(buf[0:n]))
		l.dispatch(buf[0:n], user_id)
	}
}

func (l *UserChatLogic) sendMsg(userId string, msg []byte, user_id string) {
	ctx := context.Background()
	rwLocker.RLock()
	node, ok := clientMap[userId]
	fmt.Println(node, ok, "33333333", userId, "444444", user_id)
	rwLocker.RUnlock()
	jsonMsg := Message{UserId: user_id}
	json.Unmarshal(msg, &jsonMsg)
	targetIdStr := userId
	userIdStr := jsonMsg.UserId
	jsonMsg.CreateTime = time.Now().Unix()
	r, err := l.svcCtx.RedisIm.Get(ctx, "online_"+userIdStr).Result()
	if err != nil {
		fmt.Println(err, "444444444444")
	}
	if r != "" {
		if ok {
			fmt.Println("sendMsg >>> userID: ", userId, "  msg:", string(msg))
			node.DataQueue <- msg
		}
	}
	var key string
	if userId > jsonMsg.UserId {
		key = "msg_" + userIdStr + "_" + targetIdStr
		fmt.Println(key, "&&&&&&&&&&&&&&&")
	} else {
		key = "msg_" + targetIdStr + "_" + userIdStr
		fmt.Println(key, "%%%%%%%%%%%*")
	}

	res, err := l.svcCtx.RedisIm.ZRevRange(ctx, key, 0, -1).Result()
	if err != nil {
		fmt.Println(err)
	}
	score := float64(cap(res)) + 1
	ress, e := l.svcCtx.RedisIm.ZAdd(ctx, key, &redis.Z{score, msg}).Result() //jsonMsg
	//res, e := utils.Red.Do(ctx, "zadd", key, 1, jsonMsg).Result() //备用 后续拓展 记录完整msg
	if e != nil {
		fmt.Println(e)
	}
	data := model.Contact{
		Type:     jsonMsg.Type,
		TargetId: userId,
		OwnerId:  user_id,
	}
	contact_obj, err := l.svcCtx.UserContact.FindOneByUserIdTargetId(l.ctx, user_id, userId)
	fmt.Println(err, "66666666666", contact_obj)
	if err == nil {
		if contact_obj == nil {
			_, err = l.svcCtx.UserContact.Insert(l.ctx, &data)
			if err != nil {
				fmt.Println("保存数据库失败", err)
			}
		} else {
			data = model.Contact{
				Id:        contact_obj.Id,
				TargetId:  userId,
				OwnerId:   user_id,
				CreatedAt: time.Now(),
			}
			err = l.svcCtx.UserContact.Update(l.ctx, &data)
			if err != nil {
				fmt.Println("修改数据库失败", err)
			}
		}
	}
	contact_key := fmt.Sprintf(globalkey.UserContactByUserId, user_id)
	_, e = l.svcCtx.RedisIm.ZAdd(l.ctx, contact_key, &redis.Z{float64(jsonMsg.CreateTime), userId}).Result()
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(ress, "this is a message===================")
}

func (l *UserChatLogic) sendGroupMsg(targetId string, msg []byte) {
	fmt.Println("开始群发消息")
	// userIds := SearchUserByGroupId(uint(targetId))
	// for i := 0; i < len(userIds); i++ {
	// 	//排除给自己的
	// 	if targetId != int64(userIds[i]) {
	// 		sendMsg(int64(userIds[i]), msg)
	// 	}

	// }
}
