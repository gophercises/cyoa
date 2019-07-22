package cyoa

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type handler struct {
	s Story
	t *template.Template
}

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

var tpl *template.Template

var defaultHandlerTmpl = `
<!DOCTYPE html>
<head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
            <p>{{.}}</p>
        {{end}}
    <ul>
        {{range .Options}}
            <li> <a href="/{{.Arc}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</body>`

func ParseJSON(f io.Reader) (Story, error) {
	dec := json.NewDecoder(f)
	story := make(Story)
	if err := dec.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

func NewHandler(s Story, tmpl *template.Template) http.Handler {
	if tmpl == nil {
		tmpl = tpl
	}
	return handler{s, tpl}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]
	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went Wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter Not Found.", http.StatusNotFound)

}
