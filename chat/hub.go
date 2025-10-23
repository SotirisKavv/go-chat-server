package chat

import (
	"chat-server/storage"
	"fmt"
	"sync"
)

var once sync.Once

type Hub struct {
	Rooms map[string]*Room
	Repo  storage.ChatRepository
}

var hubInstance *Hub

func GetHub() *Hub {
	if hubInstance == nil {
		once.Do(
			func() {
				hubInstance = &Hub{
					Rooms: make(map[string]*Room),
					Repo:  storage.NewChatRepository("postgres"),
				}
				go hubInstance.Repo.Run()
			},
		)
	}
	return hubInstance
}

func (h *Hub) CreateRoom(name string) *Room {
	room := NewRoom(name)
	room.Hub = h
	h.Rooms[name] = room
	go room.Run()
	return room
}

func (h *Hub) JoinRoom(client *Client, name string) {
	room, exists := h.Rooms[name]
	if !exists {
		room = h.CreateRoom(name)
	}
	client.Room = room
	room.Register <- client
}

func (h *Hub) LeaveRoom(c *Client, roomId string) {
	if room, ok := h.Rooms[roomId]; ok {
		fmt.Println("Unregister removing user ", c.Username)
		room.Unregister <- c
		if len(room.Clients) == 0 {
			delete(h.Rooms, roomId)
		}
	}
}
