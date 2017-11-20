package main

import (
	"encoding/json"
	"errors"
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

var emptyAuthor = errors.New("Author is empty")
var emptyContent = errors.New("Content is empty")

func (server *Server) addMessage(message *Message) error {
	if message.Author == "" {
		return emptyAuthor
	}
	if message.Content == "" {
		return emptyContent
	}
	server.messageChannel <- message
	return nil
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

func main() {
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
