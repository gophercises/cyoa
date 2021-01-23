package cyoa

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

const (
	intro = "intro"
)

// Handler is CYOA Story http.Handler
type Handler struct {
	story  Story
	tpl    *template.Template
	pathFn PathFn
}

// NewHandler creates new handler for provided Story
func NewHandler(opts ...HandlerOption) http.Handler {
	h := &Handler{
		story:  defaultStory,
		tpl:    defaultTpl,
		pathFn: defaultPath,
	}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

func defaultPath(r *http.Request) string {
	return strings.TrimPrefix(strings.TrimSpace(r.URL.Path), "/")
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)

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
