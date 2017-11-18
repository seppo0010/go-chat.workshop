package main

import "errors"

type Server struct {
	messages       []*Message
	messageChannel chan *Message
	channels       []chan *Message
}

func newServer() *Server {
	server := &Server{
		messages:       []*Message{},
		messageChannel: make(chan *Message),
		channels:       []chan *Message{},
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
	for _, channel := range server.channels {
		channel <- message
	}
	return nil
}

func (server *Server) getMessages() []*Message {
	return server.messages
}
func (server *Server) subscribe() chan *Message {
	channel := make(chan *Message)
	server.channels = append(server.channels, channel)
	return channel
}
func (server *Server) unsubscribe(channel chan *Message) {
	for i, thisChannel := range server.channels {
		if thisChannel == channel {
			server.channels = append(server.channels[:i], server.channels[i+1:]...)
			return
		}
	}
}

type Message struct {
	Author  string
	Content string
}
