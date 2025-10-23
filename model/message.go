package model

import "time"

type Message struct {
	SenderID  string    `json:"sender"`
	RoomID    string    `json:"room"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}
