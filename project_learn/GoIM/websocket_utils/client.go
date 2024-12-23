// websocket_utils/client.go
package websocket_utils

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

// NewClient creates a new WebSocket client
func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		Conn: conn,
		Send: make(chan []byte),
	}
}

// WritePump handles outgoing messages
func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()
	for msg := range c.Send {
		if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}
