import React, { useEffect, useRef, useState } from "react";
import { getMessageHistory, getGroupMessageHistory, sendMessageToUser, sendMessageToGroup } from "../api";
import WebSocketService from "../utils/websocket";

const Chat = () => {
    const user = JSON.parse(localStorage.getItem("user"));
    const [personalHistory, setPersonalHistory] = useState([]);
    const [groupHistory, setGroupHistory] = useState([]);
    const [message, setMessage] = useState("");
    const [receiverID, setReceiverID] = useState("");
    const [groupID, setGroupID] = useState("");
    const ws = useRef(null);

    useEffect(() => {
        ws.current = new WebSocketService(user.id);
        ws.current.connect();

        ws.current.addMessageListener((newMessage) => {
            if (newMessage.receiverID) {
                setPersonalHistory((prev) => [...prev, newMessage]);
            } else if (newMessage.groupID) {
                setGroupHistory((prev) => [...prev, newMessage]);
            }
        });

        const fetchPersonalHistory = async () => {
            const response = await getMessageHistory(user.id, receiverID);
            setPersonalHistory(response.data);
        };

        const fetchGroupHistory = async () => {
            const response = await getGroupMessageHistory(groupID);
            setGroupHistory(response.data);
        };

        if (receiverID) fetchPersonalHistory();
        if (groupID) fetchGroupHistory();

        return () => {
            ws.current.close();
        };
    }, [user.id, receiverID, groupID]);

    const sendMessage = async () => {
        const newMessage = {
            senderID: user.id,
            receiverID: receiverID || undefined,
            groupID: groupID || undefined,
            content: message,
        };

        if (receiverID) await sendMessageToUser(newMessage);
        if (groupID) await sendMessageToGroup(newMessage);

        ws.current.sendMessage(newMessage);
        setMessage("");
    };

    return (
        <div>
            <h1>Chat</h1>
            <input
                type="text"
                placeholder="Receiver ID"
                value={receiverID}
                onChange={(e) => setReceiverID(e.target.value)}
            />
            <input
                type="text"
                placeholder="Group ID"
                value={groupID}
                onChange={(e) => setGroupID(e.target.value)}
            />
            <textarea
                placeholder="Type your message..."
                value={message}
                onChange={(e) => setMessage(e.target.value)}
            />
            <button onClick={sendMessage}>Send</button>

            <h2>Personal History</h2>
            <ul>
                {personalHistory.map((msg, idx) => (
                    <li key={idx}>{msg.content}</li>
                ))}
            </ul>

            <h2>Group History</h2>
            <ul>
                {groupHistory.map((msg, idx) => (
                    <li key={idx}>{msg.content}</li>
                ))}
            </ul>
        </div>
    );
};

export default Chat;
