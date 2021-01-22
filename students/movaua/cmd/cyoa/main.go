package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

const (
	intro = "intro"
)

func main() {
	filename := flag.String("book", "goper.json", "a book JSON file name")
	port := flag.Int("port", 8080, "port to listen on for the HTTP requests")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("could not open file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	var book Book
	if err := json.NewDecoder(f).Decode(&book); err != nil {
		fmt.Printf("could not parse file: %v\n", err)
		os.Exit(1)
	}
	f.Close()

	handler := NewBookHandler(book)

	fmt.Printf("starting server on :%d...\n", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), handler); err != nil {
		fmt.Printf("could not start server on :%d: %v\n", *port, err)
		os.Exit(1)
	}
}
