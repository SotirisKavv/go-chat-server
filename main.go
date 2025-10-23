package main

import (
	"chat-server/server"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	router := server.NewRouter()

	fmt.Println("WebSocket Chat Server starting on :8080")
	fmt.Println("Open http://localhost:8080/client.html to access the chat")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
