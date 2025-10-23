package storage

import (
	"chat-server/model"
	"chat-server/utils"
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

var ctx = context.Background()

type PostgresRepository struct {
	DB           *pgx.Conn
	MessageQueue chan model.Message
	QueryBuilder utils.QueryBuilder
}

func NewPostgresRepository() (*PostgresRepository, error) {
	dbConn, err := InitDB()
	if err != nil {
		return &PostgresRepository{}, err
	}
	return &PostgresRepository{
		DB:           dbConn,
		MessageQueue: make(chan model.Message),
		QueryBuilder: utils.QueryBuilder{},
	}, nil
}

func InitDB() (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	schema := `CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
    	sender_id VARCHAR(255) NOT NULL,
    	content TEXT NOT NULL,
    	room_id VARCHAR(255) NOT NULL,
    	timestamp TIMESTAMP NOT NULL
    );`
	_, err = conn.Exec(ctx, schema)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (r *PostgresRepository) GetMessages(roomId string) ([]model.Message, error) {
	query := r.QueryBuilder.
		Select("sender_id", "content", "room_id", "timestamp").
		From("messages").
		Where(map[string]string{"room_id": "$1"}).
		OrderBy(map[string]bool{"timestamp": true}).
		Build()

	rows, err := r.DB.Query(ctx, query, roomId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []model.Message
	for rows.Next() {
		var msg model.Message
		err := rows.Scan(&msg.SenderID, &msg.Content, &msg.RoomID, &msg.Timestamp)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *PostgresRepository) SaveMessage(msg model.Message) error {
	query := r.QueryBuilder.
		InsertInto("messages").
		Fields("sender_id", "content", "room_id", "timestamp").
		Values(4).Build()

	_, err := r.DB.Exec(ctx, query, msg.SenderID, msg.Content, msg.RoomID, msg.Timestamp)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) SendMessage(msg model.Message) error {
	r.MessageQueue <- msg

	return nil
}

func (r *PostgresRepository) Run() {
	for msg := range r.MessageQueue {
		r.SaveMessage(msg)
	}
}
