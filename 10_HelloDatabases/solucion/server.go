package main

import "errors"

var emptyAuthor = errors.New("Author is empty")
var emptyContent = errors.New("Content is empty")
var channelNotFound = errors.New("Channel not found")

type Message struct {
	Author  string
	Content string
}

type Server interface {
	addMessage(message *Message) error
	getMessages() ([]*Message, error)
	subscribe() (chan *Message, error)
	unsubscribe(channel chan *Message) error
}
