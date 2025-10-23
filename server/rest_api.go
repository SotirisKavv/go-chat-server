package server

import (
	"chat-server/chat"
	"chat-server/model"
	"encoding/json"
	"net/http"
)

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	roomId := r.URL.Query().Get("r")
	messages, err := chat.GetHub().Repo.GetMessages(roomId)
	if err != nil {
		messages = []model.Message{}
	}

	historyJSON, err := json.Marshal(messages)
	if err != nil {
		http.Error(w, "Error fetching history: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(historyJSON)
}
