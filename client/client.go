package client

import (
	"bytes"
	"encoding/gob"

	"github.com/gorilla/websocket"
)

type Message struct {
	Event  string
	Buffer []byte
}

func serialize(message Message) ([]byte, error) {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(message)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

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

func SendHandler(connection *websocket.Conn, message Message) error {
	buffer, err := serialize(message)

	if err != nil {
		return err
	}

	err = connection.WriteMessage(websocket.TextMessage, buffer)

	if err != nil {
		return err
	}

	return nil
}
