package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

type Data struct {
	Message string `json:"message"`
}

func main() {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8000", Path: "/ws/1/"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer c.Close()

	// create data to send
	data := Data{Message: "hello from go"}

	// convert data to json
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("marshal error: ", err)
		return
	}
	fmt.Println(jsonData)
	// send json data
	err = c.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		fmt.Println("write error:", err)
		return
	}

	fmt.Println("Data sent to server in json format")

	// receive message
	_, message, err := c.ReadMessage()
	if err != nil {
		fmt.Println("read error:", err)
		return
	}
	fmt.Printf("received: %s\n", message)
}