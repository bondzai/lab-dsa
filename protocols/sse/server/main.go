package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/events", sseHandler)
	http.ListenAndServe(":8080", nil)
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Create a channel to send messages to the client
	messageChannel := make(chan string)

	// Close the channel when the connection is closed
	defer close(messageChannel)

	// Send a welcome message immediately
	sendMessage(w, "Welcome to SSE!")

	// Periodically send messages to the client
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-r.Context().Done():
			fmt.Println("Client connection closed.")
			return
		case <-ticker.C:
			sendMessage(w, time.Now().Format(time.RFC3339))
		}
	}
}

func sendMessage(w http.ResponseWriter, message string) {
	// Format the message in the SSE format
	fmt.Fprintf(w, "data: %s\n\n", message)
	w.(http.Flusher).Flush() // Flush the response to send the message immediately
}
