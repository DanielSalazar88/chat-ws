package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	clients map[*client]bool
	join    chan *client
	leave   chan *client
	forward chan []byte
}

func NewRoom() *room {
	r := &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}

	go r.run()
	return r
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.Receive)
		case msg := <-r.forward:
			for client := range r.clients {
				client.Receive <- msg
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (for development only)
	},
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("ServeHTTP:", err)
		return
	}
	client := &client{
		Socket:  socket,
		Receive: make(chan []byte, messageBufferSize),
		Room:    r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.Write()
	client.Read()
}
