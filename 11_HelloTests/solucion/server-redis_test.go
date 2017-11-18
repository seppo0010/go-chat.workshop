package main

import "testing"

func TestSubscribeUnsubscribeRedis(t *testing.T) {
	testSubscribeUnsubscribe(t, newServerRedis())
}

func TestSubscribePublishRedis(t *testing.T) {
	testSubscribePublish(t, newServerRedis())
}

func TestPublishGetMessagesRedis(t *testing.T) {
	testPublishGetMessages(t, newServerRedis())
}
