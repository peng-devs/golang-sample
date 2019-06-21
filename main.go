package main

import (
	"fmt"
	"net/http"
	"os"
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
	var port = ":" + os.Getenv("APP_PORT")
	fmt.Println("start server at " + port)

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/health", health)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
