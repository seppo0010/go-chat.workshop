package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Message struct {
	Author  string
	Content string
}

var messages []Message = []Message{}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		messages = append(messages, Message{
			Author:  r.PostFormValue("author"),
			Content: r.PostFormValue("content"),
		})
	} else if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(messages)
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
