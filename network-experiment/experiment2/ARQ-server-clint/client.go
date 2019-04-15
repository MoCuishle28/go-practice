package main


import (
	"bufio"
	"os"
	"net"
	"fmt"
	"strings"
	"time"
)


func main() {
	input := bufio.NewScanner(os.Stdin)
	
	fmt.Print("请输入连接IP:")
	input.Scan()

	// ip := "127.0.0.1:1200"
	ip := input.Text()

	tcpAddr, err := net.ResolveTCPAddr("tcp4", ip)		// 转换为 TCPAddr
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)		// 建立TCP连接,获得 TCPConn
	checkError(err)
	defer conn.Close()
	fmt.Println("success connection!")

	signal := make(chan int)
	closeSignal := make(chan int)
	clock := time.NewTimer(3*time.Second)

	go SendMsg(conn, input, signal, closeSignal)

	for {
		select {

		case <-signal:	// 收到ACK
			signal<-1
			clock.Reset(3*time.Second)
		case <-clock.C:	// 超时
			signal<-1
			fmt.Println("超时...")
			clock.Reset(3*time.Second)
		case <-closeSignal:	// 退出
			fmt.Println("Close...")
			return
		}
	}
}


func SendMsg(conn net.Conn, input *bufio.Scanner, signal chan int, closeSignal chan int) {
	for {
		input.Scan()
		command := input.Text()
		if command == "exi" {
			_, err := conn.Write([]byte(command))
			checkError(err)
			closeSignal<-1
			return
		} else {
			currIndex := 0
			go RecvACK(conn, signal, command, &currIndex)	// 读取数据
			go SendFrame(conn, command, signal, &currIndex)
		}
	}
}


func SendFrame(conn net.Conn, msg string, signal chan int, currIndex *int) {
	msgBytes := []byte(msg)
	sendBytes := make([]byte, 1)

	sendBytes[0] = msgBytes[*currIndex]
	conn.Write(sendBytes)
	for {
		<-signal
		if *currIndex == len(msgBytes) {
			break
		}

		sendBytes[0] = msgBytes[*currIndex]
		fmt.Println("send: ", string(msg[*currIndex]))
		conn.Write(sendBytes)
	}
}


func RecvACK(conn net.Conn, signal chan int, msg string, currIndex *int) {
	readBuff := make([]byte, 1024)
	index := 1
	for {
		size, err := conn.Read(readBuff)
		checkError(err)

		buff := string(readBuff[:size])
		index = -index + 1
		fmt.Println(strings.TrimSpace(buff), " ", index, " ", string(msg[*currIndex]))
		(*currIndex)++
		signal<-1
		if *currIndex == len(msg) {
			break
		}
		readBuff = make([]byte, 1024)
	}
}


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}