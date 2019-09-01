package main

import (
	"flag"
	"fmt"
	"github.com/gophercises/cyoa"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var port = flag.Int("port", 3000, "the port to start the CYOA web application on")
var jsonFile = flag.String("json", "gopher.json", "the JSON file with the CYOA story")
var htmlFile = flag.String("html", "", "external html template file")

func main() {
	flag.Parse()
	story := getStory()
	var options []cyoa.HandlerOption
	if *htmlFile != "" {
		tpl := getTemplateFromFile(htmlFile)
		options = append(options, cyoa.CustomTemplate(tpl))
	}
	options = append(options, cyoa.CustomChapterParser(getCustomParser()))

	start(story, options)
}

func start(story cyoa.Story, options []cyoa.HandlerOption) {
	normalHandler := cyoa.NewHandler(story)
	customHandler := cyoa.NewHandler(story, options...)

	mux := http.NewServeMux()

	mux.Handle("/", normalHandler)
	mux.Handle("/story/", customHandler)

	fmt.Printf("Server running and listening on http://localhost:%d/", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))

}

func getTemplateFromFile(file *string) *template.Template {
	tpl, err := template.ParseFiles(*file)
	if err != nil {
		panic(err)
	}
	return tpl
}

func getStory() cyoa.Story {
	file, err := os.Open(*jsonFile)
	if err != nil {
		log.Panicf("Cannot open story json file: %s\nDetails: %s", jsonFile, err)
	}
	story, err := cyoa.JsonStory(file)
	if err != nil {
		panic(err)
	}
	file.Close()
	return story
}

func getCustomParser() cyoa.ChapterParser {
	return func(r *http.Request) string {
		path := strings.TrimSpace(r.URL.Path)
		path = path[len("/story/"):]
		if path == "" {
			path = "intro"
		}
		return path
	}
}
