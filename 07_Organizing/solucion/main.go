package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var server *Server

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := server.addMessage(&Message{
			Author:  r.PostFormValue("author"),
			Content: r.PostFormValue("content"),
		})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			j, _ := json.Marshal(map[string]string{"error": err.Error()})
			w.Write(j)
			return
		}
		return
	}
	if r.Method == http.MethodGet {
		j, _ := json.Marshal(server.getMessages())
		w.Write(j)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func main() {
	server = newServer()
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
