// controllers/websocket_controller.go
package controllers

import (
	"GoIM/websocket_utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for testing; adjust for production
	},
}

// WebSocketHandler establishes a WebSocket connection
func WebSocketHandler(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
		return
	}

	client := websocket_utils.NewClient(conn)
	websocket_utils.Manager.AddClient(userID, client)

	go client.WritePump()

	// Cleanup when connection closes
	defer websocket_utils.Manager.RemoveClient(userID)
}
