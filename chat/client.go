package chat

import (
	"chat-server/model"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Room     *Room
	Send     chan *model.Message
	Username string
}

func NewClient(conn *websocket.Conn, username string) *Client {
	return &Client{
		Conn:     conn,
		Send:     make(chan *model.Message, 8), // buffered channel to always store values in buffer (sync issues)
		Username: username,
	}
}

func (c *Client) ReadMessages() {
	for {
		fmt.Println("ReadMessages waiting for message from", c.Username)
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if !isExpectedCloseError(err) {
				fmt.Println("ReadMessages error for", c.Username, ":", err)
			}
			break
		}
		if c.Room != nil {
			msg := &model.Message{
				Content:   string(message),
				SenderID:  c.Username,
				RoomID:    c.Room.Name,
				Timestamp: time.Now(),
			}
			c.Room.Hub.Repo.SendMessage(*msg)
			c.Room.Broadcast <- msg
		}
	}
}

func (c *Client) WriteMessages() {
	fmt.Println("WriteMessages started for", c.Username)
	for message := range c.Send {
		data, err := json.Marshal(message)
		if err != nil {
			fmt.Println("Error marshalling message: ", err.Error())
			continue
		}
		if err := c.Conn.WriteMessage(websocket.TextMessage, data); err != nil {
			if !isExpectedCloseError(err) {
				fmt.Println("Error writing message: ", err.Error())
			}
			break
		}
	}
	fmt.Println("WriteMessages exited for", c.Username)
}

func isExpectedCloseError(err error) bool {
	if websocket.IsCloseError(err,
		websocket.CloseGoingAway,
		websocket.CloseNormalClosure,
		1005,
	) {
		return true
	}

	return strings.Contains(err.Error(), "close sent")
}
