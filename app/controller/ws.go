package controller

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Connection struct {
	wsConn    *websocket.Conn //websocket.Conn对象
	inChan    chan []byte     //用于接收消息
	outChan   chan []byte     //用于发送消息
	closeChan chan byte       //帮助内部逻辑判断连接是否被中断
	mutex     sync.Mutex      //用于加锁
	isClose   bool            //用于避免重复关闭中的不安全因素
}

const (
	MAXMEGNUM = 1024 //最大消息数
)

//NewWS 构造一个WSConn
func NewWS(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:    wsConn,
		inChan:    make(chan []byte, MAXMEGNUM),
		outChan:   make(chan []byte, MAXMEGNUM),
		closeChan: make(chan byte, 1),
	}
	go conn.readLoop()  //执行内部读逻辑
	go conn.writeLoop() //执行内部写逻辑
	return
}

//ReadMessage 读取一条消息，即接收
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("连接已关闭。")
	}
	return
}

//WriteMessage 写一条消息，即发送
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("连接已关闭。")
	}
	return
}

func (conn *Connection) Close() {
	conn.wsConn.Close() //这个方法是线程安全的
	conn.mutex.Lock()   //上锁
	//对Chan进行处理，并标记连接已关闭
	if conn.isClose {
		close(conn.closeChan)
		conn.isClose = true
	}
	conn.mutex.Unlock() //解锁
}

func (conn *Connection) readLoop() {
	for {
		megType, data, err := conn.wsConn.ReadMessage()
		//TODO 消息类型处理
		fmt.Println(megType)
		if err != nil {
			// util.ErrHandle(err)
			conn.Close()
			return
		}
		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			conn.Close()
			return
		}
	}
}
func (conn *Connection) writeLoop() {
	var data []byte
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			conn.Close()
			return
		}
		//TODO 消息类型处理
		err := conn.wsConn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			// util.ErrHandle(err)
			conn.Close()
			return
		}
	}
}
