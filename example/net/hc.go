package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
)

var k int = 0

func main() {
	var n int64 = 4

	if len(os.Args) == 2 {
		n, _ = strconv.ParseInt(os.Args[1], 10, 32)
	}
	chs := make([]chan int, int(n))
	for i := 0; i < int(n); i++ {

		chs[i] = make(chan int)
		go Connect(i, chs[i])

	}
	for _, ch := range chs {
		<-ch
	}
	os.Exit(0)
}

func Connect(i int, ch chan int) {
	service := "163.177.65.160:80"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	k = i
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.1\r\n\r\n"))
	checkError(err)
	conn.CloseWrite()
	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))
	conn.Close()
	ch <- 1
}
func checkError(err error) {
	if err != nil {
		fmt.Println("seek err:", k)
		fmt.Fprintf(os.Stderr, "xxFatal error: %s", err.Error())
		//os.Exit(1)
	}
}
