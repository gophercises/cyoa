package main

import (
	"cyoa/server"
	"log"
	"net/http"
)

func main() {

	err := server.ServeStory("gopher.json")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))

}
