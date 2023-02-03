package decodeJsonStory

import (
	"encoding/json"
	"io"
	"os"
)

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)

	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

func ReadJsonStory(file string) (Story, error) {

	jsonFile, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	story, err := JsonStory(jsonFile)

	if err != nil {
		return nil, err
	}
	return story, nil
}
