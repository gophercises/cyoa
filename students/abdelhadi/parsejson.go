package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Arc struct {
	Title   string   `json:"title"`
	Stories []string `json:"story"`
	Options []Option `json:"options"`
}

type Story map[string]Arc

func getStories() (Story, error) {
	Stories := Story{}
	storysJson, err := os.ReadFile("gopher.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(storysJson, &Stories); err != nil {
		log.Fatal(err)
	}
	return Stories, nil
}

func getStorieByTitle(title string) (any, error) {
	Stories, err := getStories()
	if err != nil {
		log.Fatal(err)
	}

	if singleStory, ok := Stories[title]; ok {
		return singleStory, nil
	}

	return nil, fmt.Errorf("Story Not found")
}
