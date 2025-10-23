package server

import (
	"chat-server/chat"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	hub := chat.GetHub()
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Error upgrading connection: %v\n", err)
		return
	}

	roomId := r.URL.Query().Get("r")
	if roomId == "" {
		roomId = "general"
	}

	user := r.URL.Query().Get("u")
	if user == "" {
		user = "Anonymous"
	}

	fmt.Printf("New client connected: %s joining room: %s\n", user, roomId)

	client := chat.NewClient(conn, user)
	hub.JoinRoom(client, roomId)

	go client.WriteMessages()

	defer func() {
		fmt.Printf("Client %s disconnecting from room: %s\n", user, roomId)
		hub.LeaveRoom(client, roomId)
		conn.Close()
	}()

	client.ReadMessages()
}
