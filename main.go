package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[/hello] recieve request")
	w.Write([]byte("Hi, I'm Go server."))
}

func main() {
	http.HandleFunc("/hello", hello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
