package main

import (
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            log.Println(err)
            return
        }
        for {
            messageType, p, err := conn.ReadMessage()
            if err != nil {
                log.Println(err)
                return
            }
            log.Println(string(p))
            if err := conn.WriteMessage(messageType, p); err != nil {
                log.Println(err)
                return
            }
        }
    })
    log.Println("Server is running on port 8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
