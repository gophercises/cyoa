package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rsxb/cyoa"
)

var filename string

func init() {
	flag.StringVar(&filename, "f", "gopher.json", "file to read JSON story from")
}

func main() {
	flag.Parse()

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open JSON file: %s", err)
	}

	story, err := cyoa.FromJSON(f)
	if err != nil {
		log.Fatalf("Unable to parse story: %s", err)
	}
	h := cyoa.NewHandler(story)

	fmt.Println("Starting web server...")
	log.Fatal(http.ListenAndServe(":8080", h))
}
