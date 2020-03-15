package main

import (
	"net/http"

	"./story"
)

func main() {
	data, err := story.LoadJSON("./gopher.json")
	if err != nil {
		panic(err)
	}

	handle := story.GenerateHandle(data)
	err = http.ListenAndServe(":8080", handle)
	if err != nil {
		panic(err)
	}
}
