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
	Media      int64  `json:"media"`
	Url        string `json:""url`
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

var ml *UserChatLogic

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
	ml = l
	user_id := ctxdata.GetUidFromCtx(l.ctx)
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
	go sendProc(node)
	//6.完成接受逻辑
	go recvProc(node)
	//7.加入在线用户到缓存
	l.SetUserOnlineInfo("online_"+user_id, []byte(node.Addr), time.Duration(4*time.Hour))
}

func (l *UserChatLogic) SetUserOnlineInfo(key string, val []byte, timeTTL time.Duration) {
	ctx := context.Background()
	l.svcCtx.RedisIm.Set(ctx, key, val, timeTTL)
}

func sendProc(node *Node) {
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

func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := Message{}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			fmt.Println(err)
		}
		//心跳检测 msg.Media == -1 || msg.Type == 3
		if msg.Type == 3 {
			currentTime := uint64(time.Now().Unix())
			node.Heartbeat(currentTime)
		} else {
			dispatch(data)
			broadMsg(data) //todo 将消息广播到局域网
			fmt.Println("[ws] recvProc <<<<< ", string(data))
		}

	}
}

// 更新用户心跳
func (node *Node) Heartbeat(currentTime uint64) {
	node.HeartbeatTime = currentTime
}

// 后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}
	msg.CreateTime = time.Now().Unix()
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case 1: //私信
		fmt.Println("dispatch  data ===:", string(data), "-------------", msg.TargetId, msg.Content)
		sendMsg(msg.TargetId, data)
	case 2: //群发
		sendGroupMsg(msg.TargetId, data) //发送的群ID ，消息内容
		// case 4: // 心跳
		// 	node.Heartbeat()
		//case 4:
		//
	}
}

// func init() {
// 	fmt.Println("*******************************************************************************************")
// 	go udpSendProc()
// 	go udpRecvProc()
// 	fmt.Println("init goroutine ")
// 	fmt.Println("*******************************************************************************************")
// }

var udpsendChan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpsendChan <- data
}

// 完成udp数据发送协程
func udpSendProc() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(172, 20, 16, 20),
		Port: 0,
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
func udpRecvProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 0,
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
		dispatch(buf[0:n])
	}
}

func sendMsg(targetId string, msg []byte) {
	ctx := context.Background()
	rwLocker.RLock()
	node, ok := clientMap[targetId]
	rwLocker.RUnlock()
	jsonMsg := Message{}
	json.Unmarshal(msg, &jsonMsg)
	targetIdStr := targetId
	userIdStr := jsonMsg.UserId
	jsonMsg.CreateTime = time.Now().Unix()
	r, err := ml.svcCtx.RedisIm.Get(ctx, "online_"+userIdStr).Result()
	if err != nil {
		fmt.Println(err, "444444444444")
	}
	if r != "" {
		if ok {
			fmt.Println("sendMsg >>> userID: ", targetId, "  msg:", string(msg))
			node.DataQueue <- msg
		}
	}
	var key string
	if targetId > jsonMsg.UserId {
		key = "msg_" + userIdStr + "_" + targetIdStr
	} else {
		key = "msg_" + targetIdStr + "_" + userIdStr
	}
	res, err := ml.svcCtx.RedisIm.ZRevRange(ctx, key, 0, -1).Result()
	if err != nil {
		fmt.Println(err)
	}
	score := float64(cap(res)) + 1
	ress, e := ml.svcCtx.RedisIm.ZAdd(ctx, key, &redis.Z{score, msg}).Result() //jsonMsg
	//res, e := utils.Red.Do(ctx, "zadd", key, 1, jsonMsg).Result() //备用 后续拓展 记录完整msg
	if e != nil {
		fmt.Println(e)
	}
	//存入数据库逻辑
	save_db(userIdStr, targetId, jsonMsg.Type)
	save_db(targetId, userIdStr, jsonMsg.Type)
	contact_userId_key := fmt.Sprintf(globalkey.UserContactByUserId, jsonMsg.UserId)
	//添加联系列表缓存
	_, e = ml.svcCtx.RedisIm.ZAdd(ctx, contact_userId_key, &redis.Z{float64(jsonMsg.CreateTime), targetId}).Result()
	if e != nil {
		fmt.Println(e)
	}
	contact_targetId_key := fmt.Sprintf(globalkey.UserContactByUserId, targetId)
	//添加对方联系列表缓存
	_, e = ml.svcCtx.RedisIm.ZAdd(ctx, contact_targetId_key, &redis.Z{float64(jsonMsg.CreateTime), jsonMsg.UserId}).Result()
	if e != nil {
		fmt.Println(e)
	}
	user_chat_count_key := fmt.Sprintf(globalkey.UserChatCount, targetIdStr)
	ml.svcCtx.RedisIm.ZIncrBy(ctx, user_chat_count_key, 1, userIdStr)
	fmt.Println(ress, "this is a message===================")
}

func sendGroupMsg(targetId string, msg []byte) {
	fmt.Println("开始群发消息")
	// userIds := SearchUserByGroupId(uint(targetId))
	// for i := 0; i < len(userIds); i++ {
	// 	//排除给自己的
	// 	if targetId != int64(userIds[i]) {
	// 		sendMsg(int64(userIds[i]), msg)
	// 	}

	// }
}

func save_db(user_id, target_id string, ty int64) {
	data := model.Contact{
		Type:     ty,
		TargetId: target_id,
		OwnerId:  user_id,
	}
	contact_obj, err := ml.svcCtx.UserContact.FindOneByUserIdTargetId(ml.ctx, user_id, target_id)
	if err == nil {
		if contact_obj == nil {
			_, err = ml.svcCtx.UserContact.Insert(ml.ctx, &data)
			if err != nil {
				fmt.Println("保存数据库失败", err)
			}
		} else {
			data = model.Contact{
				Id:        contact_obj.Id,
				TargetId:  target_id,
				OwnerId:   user_id,
				CreatedAt: time.Now(),
			}
			err = ml.svcCtx.UserContact.Update(ml.ctx, &data)
			if err != nil {
				fmt.Println("修改数据库失败", err)
			}
		}
	}
}
