package main

import (
	"testing"
	"time"
)

func assertSubscribed(t *testing.T, server Server, expected int) {
	numSubscribed, err := server.numSubscribed()
	if err != nil {
		t.Error(err)
	}
	if numSubscribed != expected {
		t.Errorf("expected %d subscribed, got %d", expected, numSubscribed)
	}
}

func testSubscribeUnsubscribe(t *testing.T, server Server) {
	assertSubscribed(t, server, 0)
	c, err := server.subscribe()
	if err != nil {
		t.Error(err)
	}

	time.Sleep(time.Millisecond) // it might take some time to go routines to catch up
	assertSubscribed(t, server, 1)

	err = server.unsubscribe(c)
	if err != nil {
		t.Error(err)
	}

	time.Sleep(time.Millisecond) // it might take some time to go routines to catch up
	assertSubscribed(t, server, 0)
}

func testSubscribePublish(t *testing.T, server Server) {
	c, err := server.subscribe()
	if err != nil {
		t.Error(err)
	}
	time.Sleep(time.Millisecond) // it might take some time to go routines to catch up

	author, content := "author", "content"
	err = server.addMessage(&Message{Author: author, Content: content})
	if err != nil {
		t.Error(err)
	}

	message := <-c

	if message.Author != author {
		t.Errorf("expected author %s to equal %s", message.Author, author)
	}
	if message.Content != content {
		t.Errorf("expected content %s to equal %s", message.Content, content)
	}

	err = server.unsubscribe(c)
	if err != nil {
		t.Error(err)
	}
}

func testPublishGetMessages(t *testing.T, server Server) {
	now := time.Now().String()
	err := server.addMessage(&Message{Author: "author", Content: now})
	if err != nil {
		t.Error(err)
	}
	time.Sleep(time.Millisecond) // it might take some time to go routines to catch up

	messages, err := server.getMessages()
	if err != nil {
		t.Error(err)
	}
	content := messages[len(messages)-1].Content
	if content != now {
		t.Errorf("expected content to be %s, got %s", now, content)
	}
}
