package server

import (
	"cyoa/decodeJsonStory"
	"html/template"
	"net/http"
)

type Chapter decodeJsonStory.Chapter

const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		{{range .Story}}
		<div>
			<br>
			{{ . }}
			<br/>
		</div>
		{{end}}
		<br/>
		<ul>
		{{range .Options}}
		<li><a href="/{{ .Arc }}">{{ .Text }}</li>
		{{end}}
		<ul/>
	</body>
</html>`

func (h *Chapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := Chapter{
		Title:   h.Title,
		Story:   h.Story,
		Options: h.Options,
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func ServeStory(filePath string) error {
	story, error := decodeJsonStory.ReadJsonStory(filePath)

	if error != nil {
		return error
	}

	for key, value := range story {
		http.Handle("/"+key, &Chapter{
			Title:   value.Title,
			Story:   value.Story,
			Options: value.Options,
		})

	}
	return nil
}
