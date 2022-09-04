package main

import (
	"fmt"
	"log"
	"websocket-events-handler/client"

	"github.com/gorilla/websocket"
)

func receive(conn *websocket.Conn) {
	client.ReceiveHandler(conn, func(conn *websocket.Conn) {
		client.Subscribe(conn, "event", func(buffer []byte) {
			log.Printf("%s\n", buffer)
		})
	})
}

func main() {
	conn, err := client.Connect("localhost:8080")

	if err != nil {
		log.Println("Error on connecting to server: ", err)
		return
	}

	go receive(conn)

	for {
		str := ""
		fmt.Scanf("%s", &str)

		err := client.SendHandler(conn, client.Message{
			Event:  "event",
			Buffer: []byte(str),
		})

		if err != nil {
			log.Println("Error on sending message: ", err)
			return
		}
	}
}
