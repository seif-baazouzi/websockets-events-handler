package server

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Event  string
	Buffer []byte
}

func deserialize(buffer []byte) (Message, error) {
	var message Message

	b := bytes.Buffer{}
	b.Write(buffer)

	decoder := gob.NewDecoder(&b)
	err := decoder.Decode(&message)

	return message, err
}

var (
	connections = make(map[string][]*websocket.Conn)
)
var upgrader = websocket.Upgrader{}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error during connection upgradation: ", err)
		return
	}

	defer conn.Close()

	for {
		messageType, buffer, err := conn.ReadMessage()

		if err != nil {
			log.Println("Error during message reading: ", err)
			break
		}

		message, err := deserialize(buffer)

		if err != nil {
			log.Println("Error during deserialize message: ", err)
			break
		}

		if message.Event == "subscribe" {
			eventName := fmt.Sprintf("%s", message.Buffer)
			connections[eventName] = append(connections[eventName], conn)
			log.Printf("Subscribe to event %s\n", eventName)

			continue
		}

		log.Printf("%s: %s\n", message.Event, message.Buffer)

		_, foundEvent := connections[message.Event]

		if !foundEvent {
			continue
		}

		for _, conn := range connections[message.Event] {
			err = conn.WriteMessage(messageType, buffer)

			if err != nil {
				log.Println("Error during message writing:", err)
				continue
			}
		}
	}
}
