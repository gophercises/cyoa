package main

// TODO: HTML templates

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type StoryHandler struct {
	stories map[string]StoryArc // NOTE: No need for a lock because Stories is read-only
}

func (h StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:] // Get rid of the leading '/'
	if key == "" {
		key = "intro"
	}
	log.Printf("arc = %s\n", key)
	story, ok := h.stories[key]
	if !ok {
		log.Printf("story not found\n")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	storyb, err := encodeStory(story)
	if err != nil {
		// TODO OPT: Or log.Print and return 500?
		log.Panicf("encoding: %v\n", err)
	}
	if _, err := w.Write(storyb); err != nil { // NOTE: Ignoring the # of bytes written
		// TODO OPT: Or log.Print and return 500?
		log.Panicf("writing: %v\n", err)
	}
}

func encodeStory(story StoryArc) ([]byte, error) {
	storyb, err := json.MarshalIndent(story, "", "  ")
	if err != nil {
		return nil, err
	}
	return storyb, nil
}

func decodeStories(encoded []byte) (map[string]StoryArc, error) {
	stories := make(map[string]StoryArc)
	if err := json.Unmarshal(encoded, &stories); err != nil {
		return nil, err
	}
	return stories, nil
}

func main() {
	storiesb, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		log.Panicf("reading: %v\n")
	}
	var handler StoryHandler
	handler.stories, err = decodeStories(storiesb)
	if err != nil {
		log.Panicf("decoding: %v\n")
	}

	http.Handle("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
