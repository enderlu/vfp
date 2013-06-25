package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

func main() {

	service := "127.0.0.1:80"
	fmt.Println(service, "Server started.")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)

	checkError(err)

	for {

		conn, err := listener.Accept()

		if err != nil {
			continue
		}
		go Deal(conn)

	}
}

func Deal(conn net.Conn) {
	result, err := ioutil.ReadAll(conn)

	checkError(err)

	fmt.Println(string(result))

	daytime := string(result) + "你好！现在时间是：" + time.Now().String()
	conn.Write([]byte(daytime)) // don 't care about return value

	conn.Close() // we 're finished with this client
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
