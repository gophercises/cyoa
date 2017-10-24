package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gophercises/cyoa"
)

func main() {
	// Create flags for our optional variables
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	// Open the JSON file and parse the story in it.
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	// Create our custom CYOA story handler
	tpl := template.Must(template.New("").Parse(storyTmpl))
	h := cyoa.NewHandler(story,
		cyoa.WithTemplate(tpl),
		cyoa.WithPathFunc(pathFn),
	)
	// Create a ServeMux to route our requests
	mux := http.NewServeMux()
	// This story handler is using a custom function and template
	// Because we use /story/ (trailing slash) all web requests
	// whose path has the /story/ prefix will be routed here.
	mux.Handle("/story/", h)
	// This story handler is using the default functions and templates
	// Because we use / (base path) all incoming requests not
	// mapped elsewhere will be sent here.
	mux.Handle("/", cyoa.NewHandler(story))
	// Start the server using our ServeMux
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

// Updated chapter parsing function. Technically you don't
// *have* to get the story from the path (it could be a
// header or anything else) but I'm not going to rename this
// since "path" is what we used in the videos.
func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

// Slightly altered tempalte to show how this feature works
var storyTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      <ul>
      {{range .Options}}
        <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
      {{end}}
      </ul>
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FCF6FC;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #797;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: underline;
        color: #555;
      }
      a:active,
      a:hover {
        color: #222;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`
