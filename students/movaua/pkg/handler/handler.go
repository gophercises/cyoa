package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"../../pkg/model"
)

const (
	intro = "intro"
)

var (
	defaultTpl *template.Template
)

func init() {
	defaultTpl = template.Must(template.New("chapter tpl").Parse(chapterHTML))
}

// Option is a function
type Option func(h *Handler)

// WithTemplate sets the provided template
func WithTemplate(t *template.Template) Option {
	return func(h *Handler) {
		h.tpl = t
	}
}

// Handler is CYOA Story http.Handler
type Handler struct {
	story model.Story
	tpl   *template.Template
}

// New creates new handler for provided Story
func New(story model.Story, opts ...Option) http.Handler {
	h := &Handler{
		story: story,
		tpl:   defaultTpl,
	}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(strings.TrimSpace(r.URL.Path), "/")

	if path == "" {
		http.Redirect(w, r, fmt.Sprintf("/%s", intro), http.StatusFound)
		return
	}

	chaper, ok := h.story[path]
	if !ok {
		http.Error(w, "Chapter is not found.", http.StatusNotFound)
		return
	}

	if err := h.tpl.Execute(w, chaper); err != nil {
		http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		log.Printf("could not execute template: %v\n", err)
	}
}
