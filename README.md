# Websockets Events Handler

This is websockets events handler to handle events between microservices.

# How to Use

## Server

You can use the server by just specify the host and the port.

```go
server.StartServer("localhost:8080")
```

## Client

First Create a connection

```go
conn, err := client.Connect("localhost:8080")
```

Then start the receive handler, note that the ReceiveHandler function contains an infinite loop.

```go
client.ReceiveHandler(conn, func(conn *websocket.Conn) {
    ...
})
```

Then you can subscribe to events.

```go
client.Subscribe(conn, "event", func(buffer []byte) {
    log.Printf("%s\n", buffer)
})
```

You can also unsubscribe form events.

```go
client.Unsubscribe(conn, "event")
```

And sending messages to events.

```go
client.SendHandler(conn, client.Message{
    Event:  "event",
    Buffer: []byte("Message to send"),
})
```

This is an example fo client.

```go
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

		if str == "exit" {
			client.Unsubscribe(conn, "event")
		}

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
```
