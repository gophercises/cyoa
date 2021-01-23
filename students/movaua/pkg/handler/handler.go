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

// StoryHandler is CYOA Story http.Handler
type StoryHandler struct {
	story model.Story
	tpl   *template.Template
}

// New creates new handler for provided Story
func New(story model.Story) *StoryHandler {
	return &StoryHandler{
		story: story,
		tpl:   template.Must(template.New("chapter tpl").Parse(chapterHTML)),
	}
}

func (h *StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
