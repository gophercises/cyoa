package main

import (
	"Manan/cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	port := flag.Int("port", 8080, "the port to start CYOA Web Application")
	flag.Parse()
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	story, err := cyoa.ParseJSON(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story, nil)
	fmt.Printf("Started server at Port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
