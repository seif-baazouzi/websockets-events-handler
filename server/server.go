package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

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
		messageType, message, err := conn.ReadMessage()

		if err != nil {
			log.Println("Error during message reading: ", err)
			break
		}

		log.Printf("%s\n", message)

		for _, conn := range connections {
			err = conn.WriteMessage(messageType, message)

			if err != nil {
				log.Println("Error during message writing:", err)
				continue
			}
		}
	}
}
