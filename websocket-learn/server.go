package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
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
		conn *websocket.Conn
		err error
		data []byte
	)
	// 完成握手应答 第三个参数是 	responseHandler
	if conn,err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	defer conn.Close()

	// 得到 websocket.Conn 长连接
	for {
		// 能传 Text, Binary 两种类型
		if _, data, err = conn.ReadMessage(); err != nil {
			return
		}
		if err = conn.WriteMessage(websocket.TextMessage, data); err!=nil {
			return
		}
		fmt.Println(data, " string:", string(data))
	}
}