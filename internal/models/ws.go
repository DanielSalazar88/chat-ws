package models

import "github.com/gorilla/websocket"

// Clients represents a single chatting user.
type Client struct {
	// socket is the ws for thi client.
	Socket *websocket.Conn

	// receive is a channel to receive messages from other channel
	Receive chan []byte

	// room  is the room this client is chatting on.
	Room *Room
}

type Room struct {
	// Clients holds all current clients in this room.
	Clients map[*Client]bool

	// Join is a channel for clients wishing to join the room.
	Join chan *Client

	// Leave is a channel for clients wishing to leave the room.
	Leave chan *Client

	// Forward is a channel that holds incoming messages that should be forwarded to the other clients
	Forward chan []byte
}
