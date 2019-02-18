package impl

import (
	"github.com/gorilla/websocket"
	"sync"
	"errors"
)

type Connection struct {
	wsConn *websocket.Conn
	inChan chan[]byte
	outChan chan[]byte
	closeChan chan byte	// 用于在关闭连接时 把chan的阻塞也终止

	mutex sync.Mutex
	isClose bool
}


func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn: wsConn,
		inChan: make(chan []byte, 10000),
		outChan: make(chan []byte, 10000),
		closeChan: make(chan byte, 1),
	}

	// 启动读协程
	go conn.readLoop()

	// 启动写协程
	go conn.writeLoop()

	return
}


// API
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select{
		case data = <- conn.inChan:
		case <-conn.closeChan:
			err = errors.New("Connection is closed")
	}
	return
}


func (conn *Connection) WriteMessage(data []byte) (err error) {
	select{
		case conn.outChan <- data:
		case <-conn.closeChan:
			err = errors.New("Connection is closed")
	}
	return
}


func (conn *Connection) Close() {
	// Close() 是线程安全的 可以并发地调用 也可以多次调用(可重入的)
	// 但是关闭底层连接并不能关闭chan 读/写 的阻塞
	conn.wsConn.Close()

	// 所以需要 closeChan 这里 chan 只能被关闭一次
	conn.mutex.Lock()
	if !conn.isClose {
		close(conn.closeChan)
		conn.isClose = true
	}
	conn.mutex.Unlock()
}


// 读消息 内部实现
func (conn *Connection) readLoop() {
	var (
		data []byte
		err error
	)
	defer conn.Close()

	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			return
		}
		select{
			case conn.inChan <- data:
				// 如果能放进chan 则执行该分支
			case <- conn.closeChan:
				// 当closeChan关闭的时候进入 (为什么???)
				return
		}
	}
}


// 发送消息
func (conn *Connection) writeLoop() {
	var(
		data []byte
		err error
	)
	defer conn.Close()

	for {
		select {
			case data = <-conn.outChan:
			case <-conn.closeChan:
				return
		}
		// 第一个参数: 发送类型, 
		if conn.wsConn.WriteMessage(websocket.TextMessage, data); err!=nil {
			return
		}
	}
}