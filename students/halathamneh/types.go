package cyoa

import (
	"html/template"
	"net/http"
)

type Option struct {
	Text string `json:"text"`
	Chapter string `json:"arc"`
}

type Chapter struct {
	Title string `json:"title"`
	Paragraphs []string `json:"story"`
	Options []Option `json:"options"`
}

type Story map[string]Chapter


type handler struct {
	s      Story
	t      *template.Template
	chapterParser ChapterParser
}

type ChapterParser func(r *http.Request) string

type HandlerOption func(h *handler)
