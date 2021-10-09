package controller

import (
	"fmt"
	. "stencil-go/app/controller/base"
	"stencil-go/app/service"
	"sync/atomic"

	"github.com/kataras/iris/v12/websocket"
)

type WebsocketVC struct {
	BaseVC

	Service service.Todo

	*websocket.NSConn `stateless:"true"`
}

var visits uint64

func increment() uint64 {
	return atomic.AddUint64(&visits, 1)
}

func decrement() uint64 {
	return atomic.AddUint64(&visits, ^uint64(0))
}

func (vc *WebsocketVC) OnNamespaceDisconnect(msg websocket.Message) error {
	newCount := decrement()
	vc.Conn.Server().Broadcast(nil, websocket.Message{
		Namespace: msg.Namespace,
		Event:     "visitor",
		Body:      []byte(fmt.Sprintf("%d", newCount)),
	})

	return nil
}

func (vc *WebsocketVC) OnNamespaceConnected(msg websocket.Message) error {
	newCount := increment()
	vc.Conn.Server().Broadcast(nil, websocket.Message{
		Namespace: msg.Namespace,
		Event:     "visitor",
		Body:      []byte(fmt.Sprintf("%d", newCount)),
	})

	return nil
}

func (vc *WebsocketVC) Save(msg websocket.Message) error {
	id := "0"
	vc.Conn.Server().Broadcast(nil, websocket.Message{
		Namespace: msg.Namespace,
		Event:     "saved",
		To:        id,
		Body:      websocket.Marshal(vc.Service.Get(id)),
	})

	return nil
}

func (vc *WebsocketVC) OnChatTest(msg websocket.Message) error {
	ctx := websocket.GetContext(vc.Conn)
	ctx.Application().Logger().Infof("[IP: %s] [ID: %s]  broadcast to other clients the message [%s]",
		ctx.RemoteAddr(), vc, string(msg.Body))

	vc.Conn.Server().Broadcast(vc, msg)
	return nil
}
