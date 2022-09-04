package server

import (
	"bytes"
	"encoding/gob"
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

var connections []*websocket.Conn
var upgrader = websocket.Upgrader{}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error during connection upgradation: ", err)
		return
	}

	defer conn.Close()
	connections = append(connections, conn)

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

		log.Printf("%s: %s\n", message.Event, message.Buffer)

		for _, conn := range connections {
			err = conn.WriteMessage(messageType, message.Buffer)

			if err != nil {
				log.Println("Error during message writing:", err)
				continue
			}
		}
	}
}
