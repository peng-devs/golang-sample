package main

import (
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("[/hello] has bean called")
	w.Write([]byte("Hi, I'm Go server."))
}

func health(w http.ResponseWriter, r *http.Request) {
	log.Println("[/health] has been called")
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/health", health)

	done := make(chan bool)
	go http.ListenAndServe(":8080", nil)
	log.Println("Server started")
	<-done
}
