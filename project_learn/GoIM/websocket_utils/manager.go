// websocket_utils/manager.go
package websocket_utils

import (
	"sync"
)

type WebSocketManager struct {
	clients map[string]*Client // Map of userID to client
	mu      sync.RWMutex       // Mutex for thread-safe access
}

var Manager = &WebSocketManager{
	clients: make(map[string]*Client),
}

// AddClient adds a client to the manager
func (m *WebSocketManager) AddClient(userID string, client *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.clients[userID] = client
}

// RemoveClient removes a client from the manager
func (m *WebSocketManager) RemoveClient(userID string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.clients, userID)
}

// GetClient retrieves a client by userID
func (m *WebSocketManager) GetClient(userID string) (*Client, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	client, ok := m.clients[userID]
	return client, ok
}
