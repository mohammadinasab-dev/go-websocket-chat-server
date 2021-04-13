package chatserver

import "github.com/gorilla/websocket"

// UserStruct is used for sending users with socket id
type UserStruct struct {
	UserName string `json:"username"`
	// UserID   string `json:"userID"`
	//UserName string `json:"username"`
}

// SocketEventStruct struct of socket events
type SocketEventStruct struct {
	EventName    string      `json:"eventName"`
	EventPayload interface{} `json:"eventPayload"`
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub                 *Hub
	webSocketConnection *websocket.Conn
	send                chan SocketEventStruct
	username            string
	userID              string
}

// JoinDisconnectPayload will have struct for payload of join disconnect
type JoinDisconnectPayload struct {
	Users []UserStruct `json:"users"`
	// UserID string       `json:"userID"`
	UserName string `json:"username"`
}
