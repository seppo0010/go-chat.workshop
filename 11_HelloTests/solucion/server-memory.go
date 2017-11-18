package main

type ServerMemory struct {
	messages       []*Message
	messageChannel chan *Message
	channels       []chan *Message
	addChannel     chan chan *Message
	removeChannel  chan chan *Message
}

func newServerMemory() *ServerMemory {
	server := &ServerMemory{
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
						close(channel)
						server.channels = append(server.channels[:i], server.channels[i+1:]...)
						break
					}
				}
			}
		}
	}()
	return server
}

func (server *ServerMemory) addMessage(message *Message) error {
	if message.Author == "" {
		return emptyAuthor
	}
	if message.Content == "" {
		return emptyContent
	}
	server.messageChannel <- message
	return nil
}

func (server *ServerMemory) getMessages() ([]*Message, error) {
	return server.messages, nil
}
func (server *ServerMemory) subscribe() (chan *Message, error) {
	channel := make(chan *Message)
	server.addChannel <- channel
	return channel, nil
}
func (server *ServerMemory) unsubscribe(channel chan *Message) error {
	server.removeChannel <- channel
	return nil
}

func (server *ServerMemory) numSubscribed() (int, error) {
	return len(server.channels), nil
}
