package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("starting the server at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
