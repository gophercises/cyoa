package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/movaua/cyoa/pkg/cyoa"
)

func main() {
	filename := flag.String("story", "", "a story JSON file. If not provided - default story will be used")
	port := flag.Int("port", 8080, "port to listen on for the HTTP requests")
	flag.Parse()

	var opts []cyoa.HandlerOption

	if *filename != "" {
		f, err := os.Open(*filename)
		check(err)
		defer f.Close()

		story, err := cyoa.JSONStory(f)
		check(err)
		f.Close()

		opts = append(opts, cyoa.WithStory(story))
	}

	h := cyoa.NewHandler(opts...)

	fmt.Printf("starting server on :%d...\n", *port)
	check(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
