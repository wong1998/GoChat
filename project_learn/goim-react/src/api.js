// src/api.js
import axios from "axios";

const API_BASE = "http://localhost:8080/api";

export const register = (data) => axios.post(`${API_BASE}/register`, data);

export const login = (data) => axios.post(`${API_BASE}/login`, data);

export const getMessageHistory = (userID, receiverID) =>
    axios.get(`${API_BASE}/messages/user_his`, {
        params: { user_id: userID, receiver_id: receiverID },
    });

export const getGroupMessageHistory = (groupID) =>
    axios.get(`${API_BASE}/messages/group_his`, { params: { group_id: groupID } });

export const sendMessageToUser = (message) =>
    axios.post(`${API_BASE}/messages/send_to_user`, message);

export const sendMessageToGroup = (message) =>
    axios.post(`${API_BASE}/messages/send_to_group`, message);
