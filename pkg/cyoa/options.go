package cyoa

import (
	"html/template"
	"net/http"
)

// HandlerOption is a function
type HandlerOption func(h *Handler)

// WithTemplate sets the provided template
func WithTemplate(t *template.Template) HandlerOption {
	return func(h *Handler) {
		h.tpl = t
	}
}

// PathFn is function that returns a path for request
type PathFn func(r *http.Request) string

// WithPathFunc sets the provided path function
func WithPathFunc(fn PathFn) HandlerOption {
	return func(h *Handler) {
		h.pathFn = fn
	}
}

// WithStory provides a Story for the handler
func WithStory(s Story) HandlerOption {
	return func(h *Handler) {
		h.story = s
	}
}
