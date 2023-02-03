package main

import (
	"log"
	"cyoa/server"
	"net/http"
)

func main() {

	err := server.ServeStory("gopher.json")
	if err!= nil {
        log.Fatal(err)
    }
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}