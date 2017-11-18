package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var server = newServer()

type Server struct {
	messages       []*Message
	messageChannel chan *Message
}

func newServer() *Server {
	server := &Server{
		messages:       []*Message{},
		messageChannel: make(chan *Message),
	}
	go func() {
		for {
			message := <-server.messageChannel
			server.messages = append(server.messages, message)
		}
	}()
	return server
}

func (server *Server) addMessage(message *Message) {
	server.messageChannel <- message
}

func (server *Server) getMessages() []*Message {
	return server.messages
}

type Message struct {
	Author  string
	Content string
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		server.addMessage(&Message{
			Author:  r.PostFormValue("author"),
			Content: r.PostFormValue("content"),
		})
	} else if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(server.getMessages())
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
