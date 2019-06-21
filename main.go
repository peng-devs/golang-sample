package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[/hello] recieve request")
	w.Write([]byte("Hi, I'm Go server."))
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[/health] recieve request")
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/health", health)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
