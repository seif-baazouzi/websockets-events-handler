package main

import (
	"fmt"
	"log"
	"websocket-events-handler/client"

	"github.com/gorilla/websocket"
)

func receive(conn *websocket.Conn) {
	client.ReceiveHandler(conn, func(buffer []byte) {
		log.Printf("%s\n", buffer)
	})
}

func main() {
	conn, err := client.Connect("ws://localhost:8080/events")

	if err != nil {
		log.Println("Error on connecting to server: ", err)
		return
	}

	go receive(conn)

	for {
		str := ""
		fmt.Scanf("%s", &str)

		err := client.SendHandler(conn, []byte(str))
		if err != nil {
			log.Println("Error on sending message: ", err)
			return
		}
	}
}
