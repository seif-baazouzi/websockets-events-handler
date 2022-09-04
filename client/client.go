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

func deserialize(buffer []byte) (Message, error) {
	var message Message

	b := bytes.Buffer{}
	b.Write(buffer)

	decoder := gob.NewDecoder(&b)
	err := decoder.Decode(&message)

	return message, err
}

func Connect(socketUrl string) (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
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

type CallBack func([]byte)

var (
	events = make(map[string]CallBack)
)

func ReceiveHandler(conn *websocket.Conn, onStart func(*websocket.Conn)) error {
	onStart(conn)

	for {
		_, buffer, err := conn.ReadMessage()

		if err != nil {
			return err
		}

		message, err := deserialize(buffer)

		if err != nil {
			return err
		}

		callback, foundEvents := events[message.Event]

		if !foundEvents {
			continue
		}

		callback(message.Buffer)
	}
}

func Subscribe(conn *websocket.Conn, event string, callback CallBack) error {
	err := SendHandler(conn, Message{
		Event:  "subscribe",
		Buffer: []byte(event),
	})

	if err != nil {
		return err
	}

	events[event] = callback

	return nil
}
