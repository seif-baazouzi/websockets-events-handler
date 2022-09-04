package main

import (
	"log"
	"net/http"
	"websocket-events-handler/server"
)

func main() {
	http.HandleFunc("/events", server.SocketHandler)
	log.Println("Server started on port 8080")

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
