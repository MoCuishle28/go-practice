package main

import (
	"github.com/gorilla/websocket"
	"Go-practice/websocket-learn/impl"

	"fmt"
	"bufio"

	"net/http"
	"os"
	"time"
)


var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func (r *http.Request) bool {
			return true;
		},
	}
)


func recvACK(conn *impl.Connection, signal chan int, msgFrameArray []byte) {
	var recvIndex int = 0
	index := 0
	size := len(msgFrameArray)

	for {
		if _, err := conn.ReadMessage(); err != nil {
			fmt.Println("recv error!")
			return
		}
		fmt.Println("ack", index, string(msgFrameArray[recvIndex]))
		signal <- recvIndex
		recvIndex++
		index = -index + 1
		if size == recvIndex {
			break
		}
	}
}


func sendFrame(conn *impl.Connection, signal chan int, msgFrameArray []byte) {
	timer := time.NewTimer(3*time.Second)

	sendBytes := make([]byte, 1)
	endIndex := len(msgFrameArray)
	currIndex := 0

	for {
		if currIndex == endIndex{
			break
		}

		sendBytes[0] = msgFrameArray[currIndex]
		if err := conn.WriteMessage(sendBytes); err != nil {
			fmt.Println("send error!")
			return
		}
		fmt.Println("send: ", string(sendBytes[0]))

		select {

		case <-signal:	// 收到ACK
			currIndex++
			timer.Reset(3*time.Second)
		case <-timer.C:	// 超时
			fmt.Println("超时")
			timer.Reset(3*time.Second)

		}
	}

}


func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err error
		conn *impl.Connection
	)
	// defer conn.Close()	// 这里会导致空指针错误?

	// 完成握手应答 第三个参数是 	responseHandler
	if wsConn,err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	if conn,err = impl.InitConnection(wsConn); err!=nil {
		return
	}

	signal := make(chan int)
	input := bufio.NewScanner(os.Stdin)

	for {
		input.Scan()
		msg := input.Text()
		if msg == "exit" {
			break
		}
		msgFrameArray := []byte(msg)
		fmt.Println("Send Msg Frame Array: ", msgFrameArray)

		go sendFrame(conn, signal, msgFrameArray)
		go recvACK(conn, signal, msgFrameArray)
	}
	fmt.Println("close...")
}


func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}