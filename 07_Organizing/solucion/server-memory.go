package main

import "errors"

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
