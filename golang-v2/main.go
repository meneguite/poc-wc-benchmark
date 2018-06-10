package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func startWebsocket(conn *websocket.Conn) {
	defer conn.Close()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	go startWebsocket(conn)
}

func main() {
	http.HandleFunc("/", wsHandler)
	panic(http.ListenAndServe(":8080", nil))
}
