package main

import "testing"

func TestSubscribeUnsubscribeMemory(t *testing.T) {
	t.Parallel()
	testSubscribeUnsubscribe(t, newServerMemory())
}

func TestSubscribePublishMemory(t *testing.T) {
	t.Parallel()
	testSubscribePublish(t, newServerMemory())
}

func TestPublishGetMessagesMemory(t *testing.T) {
	t.Parallel()
	testPublishGetMessages(t, newServerMemory())
}
