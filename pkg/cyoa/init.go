//go:generate go-bindata -o=assets.gen.go -pkg=cyoa assets

package cyoa

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
)

var (
	defaultStory Story
	defaultTpl   *template.Template
)

func init() {
	defaultStoryJSON, err := Asset("assets/story.json")
	check(err)
	err = json.NewDecoder(bytes.NewReader(defaultStoryJSON)).Decode(&defaultStory)
	check(err)

	templateText, err := Asset("assets/chapter.html")
	check(err)
	defaultTpl = template.Must(template.New("chapter tpl").Parse(string(templateText)))
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
