package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
)

var server *Server

func handleMessagesRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := server.addMessage(&Message{
			Author:  r.PostFormValue("author"),
			Content: r.PostFormValue("content"),
		})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			j, _ := json.Marshal(map[string]string{"error": err.Error()})
			w.Write(j)
			return
		}
		return
	}
	if r.Method == http.MethodGet {
		j, _ := json.Marshal(server.getMessages())
		w.Write(j)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWebsocketRequest(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	channel := server.subscribe()
	for {
		message := <-channel
		if err := conn.WriteJSON(message); err != nil {
			break
		}
	}
	server.unsubscribe(channel)

}

func handleIndexRequest(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	server = newServer()
	http.HandleFunc("/", handleIndexRequest)
	http.HandleFunc("/messages", handleMessagesRequest)
	http.HandleFunc("/ws", handleWebsocketRequest)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
