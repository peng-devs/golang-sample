package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("[/hello] has bean called")
	w.Write([]byte("Hi, I'm Go server."))
}

func greeting(host string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("[/greet/ktor] has been called")

		resp, err := http.Get("http://" + host + "/hello")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(body)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func main() {
	host := os.Getenv("KTOR_HOST")

	http.HandleFunc("/health", health)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/greet/ktor", greeting(host))

	done := make(chan bool)
	go http.ListenAndServe(":8080", nil)
	log.Println("Server started")
	<-done
}
