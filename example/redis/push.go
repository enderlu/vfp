package main

import . "github.com/hoisie/redis"
import "runtime"
import "time"

var client Client

func init() {
	runtime.GOMAXPROCS(2)
	client.Addr = "127.0.0.1:6379"
	client.Db = 13 //第十三个工作区
}
func main() {

	sub := make(chan string, 1)
	sub <- "foo"
	messages := make(chan Message, 0)
	go client.Subscribe(sub, nil, nil, nil, messages)

	time.Sleep(10 * 1000 * 1000)
	client.Publish("foo", []byte("bar"))

	msg := <-messages
	println("received from:", msg.Channel, " message:", string(msg.Message))

	close(sub)
	close(messages)
}
