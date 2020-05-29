package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// The default HTML template mapping to each Chapter
const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Gopher adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    <p></p>
    {{range .Story}}
        <p>{{.}}</p>
    {{end}}
    <hr/>
    <ul>
        {{range .Options}}
		<li><a href="{{.Arc}}"> {{.Text}} </a></li>
		{{end}}
    </ul>
</body>
</html>
	`

type Chapters map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text,omitempty"`
	Arc  string `json:"arc,omitempty"`
}

func readfile(filename string) (Chapters, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// parse the json
	var chapters Chapters
	err = json.Unmarshal(bytes, &chapters)
	if err != nil {
		return nil, err
	}

	return chapters, nil
}

// Handler func
func displayError(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This url is not found "+r.URL.Path)

}

// Each time the user click on a link, the content of the target URL is parsed and displayed here
func fill(chapters Chapters, t *template.Template, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		path := strings.TrimSpace(r.URL.Path)
		if path == "/" || path == "" {
			path = "/intro"
		}
		if data, found := chapters[path[1:]]; found {
			err := t.Execute(w, data)
			if err != nil {
				log.Println(err)
				// http.Error(w, "Error occurred...", http.StatusNotFound)
				fallback.ServeHTTP(w, r)
			}
		} else {
			http.Error(w, "Chapter not found...", http.StatusNotFound)
			// http.Redirect(w, r, "intro", http.StatusFound)
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", displayError)

	tpl := template.Must(template.New("chapters").Parse(htmlTemplate))

	chapters, err := readfile("../../gopher.json")
	if err != nil {
		log.Fatal(err)
	}

	handler := fill(chapters, tpl, mux)

	// Start the server
	fmt.Println("Starting the server on :7894")
	http.ListenAndServe(":7894", handler)
}
