package main

import (
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8000", Path: "/ws/1/"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer c.Close()

	// send message
	// err = c.WriteMessage(websocket.TextMessage, []byte("hello, world!"))
	// if err != nil {
	// 	fmt.Println("write error:", err)
	// 	return
	// }

	// receive message
	_, message, err := c.ReadMessage()
	if err != nil {
		fmt.Println("read error:", err)
		return
	}
	fmt.Printf("received: %s\n", message)
}
