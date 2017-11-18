package main

import (
	"encoding/json"
	"sync"

	"github.com/go-redis/redis"
)

const redisChannelName = "messages"
const redisListName = "messageList"

type ServerRedis struct {
	client      *redis.Client
	stopChannel *sync.Map
}

func newServerRedis() *ServerRedis {
	return &ServerRedis{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
		stopChannel: &sync.Map{},
	}
}

func (server *ServerRedis) addMessage(message *Message) error {
	if message.Author == "" {
		return emptyAuthor
	}
	if message.Content == "" {
		return emptyContent
	}
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = server.client.Publish(redisChannelName, jsonMessage).Err()
	if err != nil {
		return err
	}
	err = server.client.RPush(redisListName, jsonMessage).Err()
	return err
}
func (server *ServerRedis) getMessages() ([]*Message, error) {
	messagesJSON, err := server.client.LRange(redisListName, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	messages := make([]*Message, len(messagesJSON))
	for i, messageJSON := range messagesJSON {
		messages[i] = &Message{}
		err = json.Unmarshal([]byte(messageJSON), messages[i])
		if err != nil {
			return nil, err
		}
	}
	return messages, nil
}
func (server *ServerRedis) subscribe() (chan *Message, error) {
	publishChannel := make(chan *Message)
	stopChannel := make(chan bool)
	server.stopChannel.Store(publishChannel, stopChannel)
	go func() {
		pubsub := server.client.Subscribe(redisChannelName)
		defer pubsub.Close()
		defer close(publishChannel)
		channel := pubsub.Channel()
		for {
			select {
			case msg := <-channel:
				message := &Message{}
				err := json.Unmarshal([]byte(msg.Payload), message)
				if err != nil {
					return
				}
				publishChannel <- message
			case _ = <-stopChannel:
				return
			}
		}
	}()
	return publishChannel, nil
}
func (server *ServerRedis) unsubscribe(channel chan *Message) error {
	stopChannel, found := server.stopChannel.Load(channel)
	if !found {
		return channelNotFound
	}
	stopChannel.(chan bool) <- true
	server.stopChannel.Delete(channel)
	return nil
}
