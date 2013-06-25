// Channel
package main

import (
	"container/list"
	"net"
)

type Channel struct {
	Conn net.Conn
}

var ChannelList = list.New()

func (channel *Channel) Close() {
	channel.Conn.Close()
}

/**消息接收**/
func (channel *Channel) Read(buffer []byte) (int, bool) {
	//读取消息长度 错误则关闭客户端链接 并从客户端列表中移除客户端信息 否则返回长度
	readLength, err := channel.Conn.Read(buffer)
	if err != nil {
		/*	fmt.Printf("Error: %v\n", err)
			fmt.Println("Shutting down client connection...\n")*/
		channel.Close()
		return 0, false
	}
	return readLength, true
}
