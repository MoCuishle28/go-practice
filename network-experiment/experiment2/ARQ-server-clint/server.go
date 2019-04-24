package main


import (
	"fmt"
	"net"
	"os"
	"time"
	"strings"

	"math/rand"
)


func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	defer listener.Close()

 	var userID int64 = 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("Accept user:", userID, "...")
		go handleClient(conn, userID)
		userID++
	}
}


func handleClient(conn net.Conn, userID int64) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	defer conn.Close()

	RecvMsg(conn, userID)
	fmt.Println("user: ", userID," close...")
}


func RecvMsg(conn net.Conn, userID int64) {
	request := make([]byte, 128)

	for {
		size, err := conn.Read(request)		// 读取客户发来的信息到 request 中
		checkError(err)

		command := string(request[:size])
		i := request[size-1]


		if len(command) >= 3 && command[:3] == "exi" {
			return
		} else {
			num := rand.Intn(100)
			if num < 30 {
				fmt.Println("数据帧 丢失/出错... ", strings.TrimSpace("Recv:"+command[:size-1]))
			} else {
				fmt.Print("ID:", userID, " ", strings.TrimSpace("Recv:"+command[:size-1]), " ", i)
				if num >= 30 && num < 50 {	// 随机休眠一段时间再ACK
					fmt.Print(" 发生迟到ACK...")
					time.Sleep(3*time.Second)
				}
				fmt.Println()
				ack := "ACK"
				if i == 0 {
					ack += "1"
				} else {
					ack += "0"
				}
				conn.Write([]byte(ack))
			}
		}
		request = make([]byte, 128)
	}
}


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}