package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

// BookHandler is CYOA http.Handler
type BookHandler struct {
	book    Book
	encoder *json.Encoder
	t       *template.Template
}

// NewBookHandler creates new handler for provided book
func NewBookHandler(book Book) *BookHandler {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	return &BookHandler{
		book:    book,
		encoder: encoder,
		t:       template.Must(template.New("webpage").Parse(chapterHTML)),
	}
}

func (h *BookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.encoder.Encode(*r.URL)

	path := strings.TrimPrefix(strings.TrimSpace(r.URL.Path), "/")

	if path == "" {
		http.Redirect(w, r, fmt.Sprintf("/%s", intro), http.StatusFound)
		return
	}

	chaper, ok := h.book[path]
	if !ok {
		http.NotFound(w, r)
		return
	}

	if err := h.t.Execute(w, chaper); err != nil {
		fmt.Printf("could not execute template: %v\n", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
