package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"Go-practice/websocket-learn/impl"
	"fmt"
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


func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}


func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err error
		data []byte
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

	// 为了演示线程安全
	go func() {
		for{
			if err := conn.WriteMessage([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(2*time.Second)	// 每秒发一个心跳消息给客户端
		}
	}()

	for {
		if data,err = conn.ReadMessage(); err != nil {
			return
		}
		if err = conn.WriteMessage(data); err != nil {
			return
		}
		fmt.Println("data:", string(data))
	}
}