package controllers

import (
	"GoIM/config"
	"GoIM/models"
	"GoIM/websocket_utils"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetMessageHistory retrieves message history between two users
func GetMessageHistory(c *gin.Context) {
	var messages []models.Message
	userID := c.Query("user_id")
	receiverID := c.Query("receiver_id")

	if userID == "" || receiverID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id and receiver_id are required"})
		return
	}

	// Query messages from the database
	if err := config.DB.Where(
		"(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
		userID, receiverID, receiverID, userID,
	).Order("timestamp asc").Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

// GetGroupMessageHistory retrieves message history for a group
func GetGroupMessageHistory(c *gin.Context) {
	var messages []models.GroupMessage
	groupID := c.Query("group_id")

	if groupID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "group_id is required"})
		return
	}

	// Query group messages from the database
	if err := config.DB.Where("group_id = ?", groupID).Order("timestamp asc").Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve group messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

// SendMessageToUser handles sending a message to a specific user
func SendMessageToUser(c *gin.Context) {
	var message models.Message

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	message.Timestamp = time.Now().Unix()

	if err := config.DB.Create(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save message"})
		return
	}
	// Attempt WebSocket push
	if err := pushToReceiver(message); err != nil {
		// Log WebSocket error but return success to client
		log.Printf("WebSocket push failed for receiver %d: %v", message.ReceiverID, err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})

}

// SendMessageToGroup handles sending a message to a group
func SendMessageToGroup(c *gin.Context) {
	var groupMessage models.GroupMessage

	if err := c.ShouldBindJSON(&groupMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	groupMessage.Timestamp = time.Now().Unix()

	if err := config.DB.Create(&groupMessage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save group message"})
		return
	}
	// Attempt WebSocket broadcast
	if err := broadcastToGroup(groupMessage); err != nil {
		// Log WebSocket error but return success to client
		log.Printf("WebSocket broadcast failed for group %d: %v", groupMessage.GroupID, err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Group message sent successfully"})
}

// pushToReceiver attempts to push a message via WebSocket to the receiver
func pushToReceiver(message models.Message) error {
	receiverIDStr := fmt.Sprintf("%d", message.ReceiverID)
	receiverClient, online := websocket_utils.Manager.GetClient(receiverIDStr)
	if online {
		receiverClient.Send <- []byte(message.Content)
		return nil
	}
	return fmt.Errorf("receiver %d is not online", message.ReceiverID)
}

// broadcastToGroup attempts to broadcast a message via WebSocket to all online group members
func broadcastToGroup(groupMessage models.GroupMessage) error {
	var groupMembers []models.GroupMember

	// Fetch all members of the group
	if err := config.DB.Where("group_id = ?", groupMessage.GroupID).Find(&groupMembers).Error; err != nil {
		return fmt.Errorf("failed to fetch group members: %v", err)
	}

	// Push the message to all online group members
	for _, member := range groupMembers {
		memberIDStr := fmt.Sprintf("%d", member.UserID)
		client, online := websocket_utils.Manager.GetClient(memberIDStr)
		if online {
			client.Send <- []byte(groupMessage.Content)
		}
	}

	return nil
}
