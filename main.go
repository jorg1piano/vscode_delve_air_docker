package main

import (
	"fmt"
	"log"
	"net/http"

	"example/pages"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := "Hello, World!!"
		pages.HelloWorld(msg).Render(r.Context(), w)
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
