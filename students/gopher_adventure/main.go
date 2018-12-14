package main

import (
	"bytes"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// Charlie's code
type Story map[string]Chapter
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

//// end of Charlie's code

var (
	story Story
	tpl   *template.Template
)

func main() {
	//var c interface{}

	file, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		log.Fatal(err)
	}
	r := bytes.NewReader(file)
	d := json.NewDecoder(r)
	//fmt.Printf("File contents: %s", content)
	if err := d.Decode(&story); err != nil {
		log.Fatal(err)
	}

	spew.Dump(story)

	mux := http.NewServeMux()
	mux.HandleFunc("/", start)
	//mux.HandleFunc("/story", story)
	fs := http.FileServer(http.Dir("static"))
	//mux.Handle("/static/", fs)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", mux)
}

func start(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	var key string

	if val, ok := keys["key"]; ok {
		if len(val) > 0 {
			key = val[0]
			tpl.ExecuteTemplate(w, "index", story[key])
			return
		}
	}
	tpl.ExecuteTemplate(w, "index", story["intro"])
	return
}

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}
