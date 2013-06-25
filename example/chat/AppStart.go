// AppStart
package main

import (
	"encoding/json"
	"fmt"
	"github.com/enderlu/vfp/example/chat/ace"
	"net"
)

var index = NewAtomicInteger()

func main() {
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Println("listener err")
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		fmt.Println("客户端连接")
		go clientConn(conn)
	}

}

func clientConn(conn net.Conn) {
	channel := &Channel{conn}
	go clientRead(channel)
}

func clientRead(channel *Channel) {
	buffer := make([]byte, 1024)
	buff := ace.NewBuffer()
	for bytelength, readSuccess := channel.Read(buffer); readSuccess; bytelength, readSuccess = channel.Read(buffer) {
		buff.WriteBytes(buffer[0:bytelength])
		/*length := buff.ReadInt()
		if length > buff.Length()-4 {
			buff.Reset()
			continue
		}*/
		fmt.Println("Length:", buff.Length())
		c := buff.ReadInt()
		m := buff.ReadBytes()
		go process(channel.Conn, SocketModel{c, m})
		buff.Clear()
	}
}

func process(conn net.Conn, model SocketModel) {
	switch model.Command {
	case REG_CREQ: //注册连接
		reg(conn, string(model.Message))
		break
	case CHAT_CREQ: //私聊
		chat(conn, model.Message)
		break
	case ROOM_CHAT_CREQ: //群发
		RoomChat(conn, model.Message)
		break
	}
}

/*传输协议*/
const (
	REG_CREQ       = 0
	CHAT_CREQ      = 1
	ROOM_CHAT_CREQ = 2
	REG_SRES       = 3
	REG_BRO        = 4
	CHAT_SRES      = 5
	ROOM_CHAT_BRO  = 6
)

type ChatModel struct {
	UserId  int
	message string
}

func RoomChat(conn net.Conn, message []byte) {
	var chatModel ChatModel
	err := json.Unmarshal(message, &chatModel)
	if err != nil {
		return
	}
	userId := UserMap.Get(conn)
	chatModel.UserId = userId
	exbrocast(userId, ROOM_CHAT_BRO, chatModel)
}

/*私聊*/
func chat(conn net.Conn, message []byte) {
	userId := UserMap.Get(conn)

	var chatModel ChatModel
	err := json.Unmarshal(message, &chatModel)
	if err != nil {
		return
	}
	targetId := chatModel.UserId
	info := UserInfoMap.Get(targetId)
	wConn := info.conn
	chatModel.UserId = userId
	write(wConn, CHAT_SRES, chatModel)
}

/*输入名字 进入房间---产生绑定*/
func reg(conn net.Conn, name string) {
	id := index.AddAndGet()
	UserMap.Put(conn, id)
	info := userInfo{id, name, conn}
	UserInfoMap.Put(id, info)
	fmt.Println(id, name, info)
	exbrocast(id, REG_BRO, info)
	write(conn, REG_SRES, getUsers())

}

/*发送消息*/
func write(conn net.Conn, command int, message interface{}) {
	w, err := json.Marshal(message)
	if err != nil {
		return
	}
	buff := ace.NewBuffer()
	buff.WriteInt(command)
	buff.WriteBytes(w)
	conn.Write(buff.Bytes())
}

/*获取房间所有用户--包括自己*/
func getUsers() []userInfo {
	var users []userInfo
	for _, v := range UserInfoMap.M {
		users = append(users, v)
	}
	return users
}

/*群发*/
func brocast(command int, message interface{}) {
	w, err := json.Marshal(message)
	if err != nil {
		return
	}
	buff := ace.NewBuffer()
	buff.WriteInt(command)
	buff.WriteBytes(w)
	m := UserMap.GetMap()
	for k, _ := range m {
		k.Write(buff.Bytes())
	}
}

/*排除指定用户群发*/
func exbrocast(exid int, command int, message interface{}) {
	w, err := json.Marshal(message)
	if err != nil {
		return
	}
	buff := ace.NewBuffer()
	buff.WriteInt(command)
	buff.WriteBytes(w)
	m := UserMap.GetMap()
	for k, v := range m {
		if exid == v {
			continue
		}
		k.Write(buff.Bytes())
	}
}
