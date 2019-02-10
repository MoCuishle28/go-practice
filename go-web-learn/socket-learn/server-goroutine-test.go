package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"strconv"
	"strings"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	request := make([]byte, 128) // set maxium request length to 128B to prevent flood attack
	defer conn.Close()  		// close connection before exit

	for {
		read_len, err := conn.Read(request)		// 读取客户发来的信息到 request 中
		if err != nil {
			return
		}
		fmt.Println("read_len:", read_len)

		if read_len == 0 {
			break 		// connection already closed by client
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))		// 似乎要在客户端接收才能继续执行(否则会阻塞?) 不用啊
			fmt.Println("receive:",string(request)," in timestamp")
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))		// 似乎要在客户端接收才能继续执行(否则会阻塞?) 不用啊
			fmt.Println("receive:",string(request), "in else")
		}

		request = make([]byte, 128) // clear last read content
		fmt.Println("current conn end")
		break	// 不跳出循环就需要一直从客户端发送信息过来
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}