package storage

import (
	"chat-server/model"
)

type MemoryRepository struct {
	Messages     map[string][]model.Message
	MessageQueue chan model.Message
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		Messages:     make(map[string][]model.Message),
		MessageQueue: make(chan model.Message, 8),
	}
}

func (r *MemoryRepository) Run() {
	for message := range r.MessageQueue {
		r.SaveMessage(message)
	}
}

func (r *MemoryRepository) GetMessages(roomId string) ([]model.Message, error) {
	roomHistory, ok := r.Messages[roomId]
	if !ok {
		return []model.Message{}, nil // Return empty slice instead of error for better UX
	}

	return roomHistory, nil
}

func (r *MemoryRepository) SaveMessage(msg model.Message) error {
	roomId := msg.RoomID
	r.Messages[roomId] = append(r.Messages[roomId], msg)
	return nil
}

func (r *MemoryRepository) SendMessage(msg model.Message) error {
	r.MessageQueue <- msg
	return nil
}
