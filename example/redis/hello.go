package main

import . "github.com/hoisie/redis"
import "runtime"

var client Client

func init() {
	runtime.GOMAXPROCS(2)
	client.Addr = "127.0.0.1:6379"
	client.Db = 13 //第十三个工作区
}

func main() {

	var key = "hello"
	client.Set(key, []byte("world"))
	val, _ := client.Get("hello")
	println(key, string(val))
}
