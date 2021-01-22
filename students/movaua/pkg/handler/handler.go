package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"../../pkg/model"
)

const (
	intro = "intro"
)

// StoryHandler is CYOA Story http.Handler
type StoryHandler struct {
	Story model.Story
	tpl   *template.Template
}

// New creates new handler for provided Story
func New(Story model.Story) *StoryHandler {
	return &StoryHandler{
		Story: Story,
		tpl:   template.Must(template.New("webpage").Parse(chapterHTML)),
	}
}

func (h *StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(strings.TrimSpace(r.URL.Path), "/")

	if path == "" {
		http.Redirect(w, r, fmt.Sprintf("/%s", intro), http.StatusFound)
		return
	}

	chaper, ok := h.Story[path]
	if !ok {
		http.NotFound(w, r)
		return
	}

	if err := h.tpl.Execute(w, chaper); err != nil {
		fmt.Printf("could not execute template: %v\n", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
