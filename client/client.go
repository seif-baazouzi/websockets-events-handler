package client

import (
	"github.com/gorilla/websocket"
)

func Connect(socketUrl string) (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func ReceiveHandler(connection *websocket.Conn, callback func([]byte)) error {
	for {
		_, buffer, err := connection.ReadMessage()

		if err != nil {
			return err
		}

		callback(buffer)
	}
}

func SendHandler(connection *websocket.Conn, buffer []byte) error {
	err := connection.WriteMessage(websocket.TextMessage, buffer)

	if err != nil {
		return err
	}

	return nil
}
