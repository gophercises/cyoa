package cyoa

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

// default HTML template
var tmpl *template.Template

func init() {
	filename := "story.tpl"

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	s := string(bytes)

	tmpl = template.Must(template.New("").Parse(s))
}

type handler struct {
	story     Story
	template  *template.Template
	parsePath func(r *http.Request) string
}

// parsePath gets the chapter title from a request's URL path.
func parsePath(r *http.Request) string {
	path := r.URL.Path
	if path == "" || path == "/" {
		path = "/intro"
	}
	return path[1:]
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.parsePath(r)
	if chapter, ok := h.story[path]; ok {
		err := h.template.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)
}

// HandlerOption configures Story handlers.
type HandlerOption func(h *handler)

// WithTemplate applies an html.Template to the returned handler.
func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.template = t
	}
}

// WithParser applies a custom URL parser to the returned handler.
func WithParser(pathParser func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.parsePath = pathParser
	}
}

// NewHandler returns an http.Handler that parses story templates.
func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s, tmpl, parsePath}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

// Story is a Choose-Your-Own-Adventure plotline.
type Story map[string]Chapter

// Chapter is a section of a story.
type Chapter struct {
	Title      string   `json:"title,omitempty"`
	Paragraphs []string `json:"story,omitempty"`
	Options    []Option `json:"options,omitempty"`
}

// Option is a choice presented to the user.
type Option struct {
	Text    string `json:"text,omitempty"`
	Chapter string `json:"arc,omitempty"`
}

// FromJSON converts from JSON to Story.
func FromJSON(r io.Reader) (Story, error) {
	var story Story
	d := json.NewDecoder(r)
	if err := d.Decode(&story); err != nil {
		return nil, fmt.Errorf("FromJSON: %s", err)
	}
	return story, nil
}
