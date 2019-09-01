package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

var tpl *template.Template

func init() {
	tpl, _ = template.New("").Parse(defaultHtmlTmplate)
}

func JsonStory(reader io.Reader) (Story, error) {
	decoder := json.NewDecoder(reader)
	var story Story
	err := decoder.Decode(&story)
	if err != nil {
		return nil, err
	}
	return story, nil
}

func CustomChapterParser(parser ChapterParser) HandlerOption {
	return func(h *handler) {
		h.chapterParser = parser
	}
}

func CustomTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s, tpl, defaultChapterParser}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

func defaultChapterParser(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	return path[1:]
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.chapterParser(r)
	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Print(err)
			http.Error(w, "Semething went wrong!", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found", http.StatusNotFound)
}


var defaultHtmlTmplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Create your own adventure</title>
    <style>
        html, body {
            margin: 0;
            padding: 0;
            background-color: #efefef;
        }

        .container {
            min-width: 100vw;
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
        }

        .card {
            padding: 16px;
            background-color: #fff;
            box-shadow: 0 0 0 5px rgba(150, 150, 150, .5);
            min-width: 550px;
            max-width: 1000px;
        }
    </style>
</head>
<body>
<div class="container">
    <div class="card">
        <h1>{{.Title}}</h1>
        <p>{{range .Paragraphs}}{{.}}{{end}}</p>
        <ul>
            {{range .Options}}
            <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
            {{end}}
        </ul>
    </div>
</div>
</body>
</html>
`