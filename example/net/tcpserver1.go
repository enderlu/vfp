package main

import (
	"log"
	"net"
	"runtime"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":9988")
	check_error(err)

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)

	check_error(err)

	for i := 0; i < 100; i++ {

		go handle_tcp_accept(tcpListener)

	}

	{
	}

}

func handle_tcp_accept(tcpListener *net.TCPListener) {

	for {

		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {

			log.Println("tcp accept failed!")

			continue

		} else {

			connChan := make(chan []byte)

			go write_tcp_conn(tcpConn, connChan)

			go read_tcp_conn(tcpConn, connChan)

		}

	}

}

func read_tcp_conn(tcpConn *net.TCPConn, connChan chan []byte) {

	buffer := make([]byte, 2048)

	tcpConn.SetReadBuffer(2048)

	for {

		n, err := tcpConn.Read(buffer[0:])

		if err != nil {

			log.Println("one tcp connection read function failed!")

			log.Println("one tcp connection close now!")

			tcpConn.Close()

			runtime.Goexit()

		} else {

			connChan <- buffer[0 : n-1]

		}

	}

}

func write_tcp_conn(tcpConn *net.TCPConn, connChan chan []byte) {

	for {

		msg := <-connChan

		log.Println(string(msg))

		tcpConn.Write([]byte(msg)[0 : len(msg)+1])

	}

}

func check_error(err error) {

	if err != nil {

		log.Printf("Fatal error : ％s", err.Error())

	}

}
