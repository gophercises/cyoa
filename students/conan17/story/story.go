package story

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Chapter story chapter
type Chapter struct {
	Title      string    `json:"title"`
	Paragraphs []string  `json:"story"`
	Options    []Options `json"options"`
}

// Options chapter options
type Options struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

// LoadJSON Load json file to generate story set
func LoadJSON(path string) (map[string]Chapter, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	body, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	story := make(map[string]Chapter)
	err = json.Unmarshal(body, &story)
	if err != nil {
		return nil, err
	}
	return story, nil
}

// GenerateHandle import stories and generate handler
func GenerateHandle(story map[string]Chapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimLeft(r.URL.Path, "/")
		if path == "" {
			path = "intro"
		}
		if _, ok := story[path]; !ok {
			return
		}

		body := bytes.Buffer{}
		defalutTemplate.Execute(&body, story[path])
		w.Write(body.Bytes())
	}
}

var defalutTemplate *template.Template

func init() {
	defalutTemplate, _ = template.New("").Parse(defaultHTML)
}

const defaultHTML = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      {{if .Options}}
        <ul>
        {{range .Options}}
          <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
      {{else}}
        <h3>The End</h3>
      {{end}}
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`