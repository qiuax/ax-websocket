package lib

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type WSConnect struct {
	Conn *websocket.Conn

	inChan    chan []byte
	outChan   chan []byte
	mt        sync.Mutex
	closeChan chan []byte
	isClose   bool
}

func NewConn(conn *websocket.Conn) *WSConnect {
	wsconn := &WSConnect{
		Conn:    conn,
		inChan:  make(chan []byte, 1000),
		outChan: make(chan []byte, 1000),

		closeChan: make(chan []byte, 1),
	}

	// 启动读协程
	go wsconn.readWS()

	// 启动写协程
	go wsconn.writeWS()

	return wsconn
}

// ReadMessage 读取消息
func (conn *WSConnect) ReadMessage() (data []byte, err error) {

	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("连接已经关闭")

	}
	return
}

// WriteMessage 发送消息
func (conn *WSConnect) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("连接已经关闭")
	}
	return
}

// ReadLoop 从websocket中读取消息
func (conn WSConnect) readWS() {

	for {
		_, data, err := conn.Conn.ReadMessage()
		if err != nil {
			goto ERR
		}
		select {
		case conn.inChan <- data:
			// 当closeChan关闭的时候会执行下面代码
		case <-conn.closeChan:
			goto ERR
		}
	}

ERR:
	conn.Close()
}

func (conn *WSConnect) writeWS() {
	var data []byte
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}
		if err := conn.Conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()
}

// Close 关闭连接
func (conn *WSConnect) Close() {
	conn.Conn.Close()
	// 关闭closeChan,chan 只能关闭一次的
	conn.mt.Lock()
	if !conn.isClose {
		close(conn.closeChan)
		conn.isClose = true
	}
	conn.mt.Unlock()
}
