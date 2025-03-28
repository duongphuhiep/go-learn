package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

/*
WebSocket is a communication protocol that provides full-duplex communication channels over a single TCP connection.
It is commonly used in web applications to send and receive messages in real-time.
*/
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		messageStr := string(message)
		slog.Info("Received message", "message", messageStr, "type", messageType)

		/*messageType can have the following value:  websocket.TextMessage (value 1), websocket.BinaryMessage (value 2), websocket.CloseMessage (value 8), websocket.PingMessage (value 9), websocket.PongMessage (value 10) */
		if messageStr == "hello" {
			conn.WriteMessage(messageType, []byte("hello back"))
			conn.WriteMessage(messageType, []byte("my name is websocket"))
		}
	}
}

/*
Server-Sent Events (SSE) is a standard describing how servers can initiate data transmission towards the client once
an initial client connection has been established. This uni-direction communication are commonly used to send real-time
updates to a web application.
*/
func handleSSE(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Ensure response flushing
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// Send periodic messages
	for {
		fmt.Fprintf(w, "data: time on server is %s\n\n", time.Now().Format(time.RFC3339))
		flusher.Flush()
		time.Sleep(2 * time.Second)
	}
}

func main() {
	// Serve static files (HTML, JS, CSS, etc.) from the current directory
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// WebSocket endpoint
	http.HandleFunc("/ws", handleWebSocket)
	http.HandleFunc("/events", handleSSE)

	fmt.Println("Server started at http://localhost:8080")
	fmt.Println("WebSocket endpoint: ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
