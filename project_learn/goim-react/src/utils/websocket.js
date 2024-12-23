// src/utils/websocket.js
class WebSocketService {
    constructor(userID) {
        this.userID = userID;
        this.ws = null;
        this.listeners = [];
    }

    connect() {
        this.ws = new WebSocket(`ws://localhost:8080/ws?user_id=${this.userID}`);

        this.ws.onopen = () => {
            console.log("WebSocket connected");
        };

        this.ws.onmessage = (event) => {
            const message = JSON.parse(event.data);
            this.listeners.forEach((listener) => listener(message));
        };

        this.ws.onclose = () => {
            console.log("WebSocket disconnected");
        };

        this.ws.onerror = (error) => {
            console.error("WebSocket error:", error);
        };
    }

    addMessageListener(callback) {
        this.listeners.push(callback);
    }

    removeMessageListener(callback) {
        this.listeners = this.listeners.filter((listener) => listener !== callback);
    }

    sendMessage(message) {
        if (this.ws && this.ws.readyState === WebSocket.OPEN) {
            this.ws.send(JSON.stringify(message));
        }
    }

    close() {
        if (this.ws) {
            this.ws.close();
        }
    }
}

export default WebSocketService;
