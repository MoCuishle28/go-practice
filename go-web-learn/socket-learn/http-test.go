package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	// 向 os.Args[1] 服务器地址 发送一个 http 请求 
	// 比如：183.232.231.172:80 百度服务器 (要带端口号,web服务一般是80)
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)		// 转换为 TCPAddr
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)			// 建立TCP连接,获得 TCPConn
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))  // 发送 http 请求
	checkError(err)

	result, err := ioutil.ReadAll(conn)		// 读取返回内容
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}