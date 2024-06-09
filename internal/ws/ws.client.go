package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type client struct {
	Socket  *websocket.Conn
	Receive chan []byte
	Room    *room
}

func (c *client) Read() {
	defer c.Socket.Close()
	for {
		x, msg, err := c.Socket.ReadMessage()
		log.Println(x, msg)
		if err != nil {
			return
		}
		c.Room.forward <- msg
	}
}

func (c *client) Write() {
	defer c.Socket.Close()
	for msg := range c.Receive {
		err := c.Socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
