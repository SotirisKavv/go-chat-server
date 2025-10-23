package chat

import (
	"chat-server/model"
	"fmt"
	"time"
)

type Room struct {
	Hub        *Hub
	Name       string
	Clients    map[*Client]bool
	Broadcast  chan *model.Message
	Register   chan *Client
	Unregister chan *Client
}

func NewRoom(name string) *Room {
	return &Room{
		Name:       name,
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan *model.Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.Register:
			r.Clients[client] = true
			fmt.Printf("Registered client: %s. Total clients: %d\n", client.Username, len(r.Clients))
			msg := &model.Message{
				Content:   fmt.Sprintf("User %s entered the chat\n", client.Username),
				SenderID:  "SYSTEM",
				RoomID:    r.Name,
				Timestamp: time.Now(),
			}
			r.BroadcastSystemMessage(msg)
		case client := <-r.Unregister:
			if _, ok := r.Clients[client]; ok {
				msg := &model.Message{
					Content:   fmt.Sprintf("User %s left the chat\n", client.Username),
					SenderID:  "SYSTEM",
					RoomID:    r.Name,
					Timestamp: time.Now(),
				}
				r.BroadcastSystemMessage(msg)
				delete(r.Clients, client)
				close(client.Send)
			}
		case message := <-r.Broadcast:
			for client := range r.Clients {
				select {
				case client.Send <- message:
					fmt.Printf("Broadcasting message: %+v\n", message)
				default:
					fmt.Printf("Removing client (Send blocked): %s\n", client.Username)
					close(client.Send)
					delete(r.Clients, client)
				}
			}
		}
	}
}

func (r *Room) BroadcastSystemMessage(msg *model.Message) {
	for client := range r.Clients {
		select {
		case client.Send <- msg:
		default:
			fmt.Printf("Error during broadcast System")
			close(client.Send)
			delete(r.Clients, client)
		}
	}
}
