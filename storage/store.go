package storage

import (
	"chat-server/model"
	"fmt"
)

type ChatRepository interface {
	GetMessages(roomId string) ([]model.Message, error)
	SaveMessage(msg model.Message) error
	SendMessage(msg model.Message) error
	Run()
}

func NewChatRepository(chatType string) ChatRepository {
	switch chatType {
	case "memo":
		return NewMemoryRepository()
	case "postgres":
		repo, err := NewPostgresRepository()
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return nil
		}
		return repo
	default:
		return nil
	}
}
