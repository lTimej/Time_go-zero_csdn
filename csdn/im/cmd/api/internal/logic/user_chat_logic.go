package logic

import (
	"context"
	"encoding/json"
	"fmt"
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
	user_id := "1391913700415242240"
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
	go l.recvProc(node)

	go l.udpSendProc()
	go l.udpRecvProc()
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

func (l *UserChatLogic) recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := model.Message{}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			fmt.Println(err)
		}
		//心跳检测 msg.Media == -1 || msg.Type == 3
		if msg.Type == 3 {
			currentTime := uint64(time.Now().Unix())
			node.Heartbeat(currentTime)
		} else {
			l.dispatch(data)
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
func (l *UserChatLogic) dispatch(data []byte) {
	fmt.Println(string(data), "===========")
	msg := model.Message{}
	msg.CreateTime = time.Now().Unix()
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msg, "===========")
	switch msg.Type {
	case 1: //私信
		fmt.Println("dispatch  data ===:", string(data))
		l.sendMsg(msg.TargetId, data)
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
func (l *UserChatLogic) udpRecvProc() {
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
		l.dispatch(buf[0:n])
	}
}

func (l *UserChatLogic) sendMsg(userId string, msg []byte) {

	rwLocker.RLock()
	node, ok := clientMap[userId]
	rwLocker.RUnlock()
	jsonMsg := model.Message{}
	json.Unmarshal(msg, &jsonMsg)
	targetIdStr := userId
	userIdStr := jsonMsg.UserId
	jsonMsg.CreateTime = time.Now().Unix()
	r, err := l.svcCtx.RedisIm.Get(l.ctx, "online_"+userIdStr).Result()
	if err != nil {
		fmt.Println(err)
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
	} else {
		key = "msg_" + targetIdStr + "_" + userIdStr
	}
	res, err := l.svcCtx.RedisIm.ZRevRange(l.ctx, key, 0, -1).Result()
	if err != nil {
		fmt.Println(err)
	}
	score := float64(cap(res)) + 1
	ress, e := l.svcCtx.RedisIm.ZAdd(l.ctx, key, &redis.Z{score, msg}).Result() //jsonMsg
	//res, e := utils.Red.Do(ctx, "zadd", key, 1, jsonMsg).Result() //备用 后续拓展 记录完整msg
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(ress)
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
