package main

import "errors"

type Server struct {
	messages       []*Message
	messageChannel chan *Message
	channels       []chan *Message
	addChannel     chan chan *Message
	removeChannel  chan chan *Message
}

func newServer() *Server {
	server := &Server{
		messages:       []*Message{},
		messageChannel: make(chan *Message),
		addChannel:     make(chan chan *Message),
		removeChannel:  make(chan chan *Message),
		channels:       []chan *Message{},
	}
	go func() {
		for {
			select {
			case message := <-server.messageChannel:
				server.messages = append(server.messages, message)
				for _, channel := range server.channels {
					channel <- message
				}
			case channel := <-server.addChannel:
				server.channels = append(server.channels, channel)
			case channel := <-server.removeChannel:
				for i, thisChannel := range server.channels {
					if thisChannel == channel {
						server.channels = append(server.channels[:i], server.channels[i+1:]...)
						return
					}
				}
			}
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
func (server *Server) subscribe() chan *Message {
	channel := make(chan *Message)
	server.addChannel <- channel
	return channel
}
func (server *Server) unsubscribe(channel chan *Message) {
	server.removeChannel <- channel
}

type Message struct {
	Author  string
	Content string
}
