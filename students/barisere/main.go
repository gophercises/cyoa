package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// StoryArc is a path in a story
type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func (s StoryArc) explainStoryArc() string {
	return fmt.Sprintf("%s\n\n\t%s\n", s.Title, strings.Join(s.Story, "\n\t"))
}

// Story is an adventure with story arcs
type Story map[string]StoryArc

func parseStory(storyMap []byte) (story Story) {
	if err := json.Unmarshal(storyMap, &story); err != nil {
		log.Fatalf("Error parsing story map: %s", err)
	}
	return
}

func traverseArcs(story Story, arc string) string {
	chosenArc := story[arc]
	fmt.Println(chosenArc.explainStoryArc())
	if len(chosenArc.Options) > 0 {
		fmt.Println("Choose an option:")
		for index, option := range chosenArc.Options {
			fmt.Fprintf(os.Stdout, "\n%d %s: %s\n", index+1, option.Arc, option.Text)
		}
		var choice int
		if _, err := fmt.Fscanf(os.Stdin, "%d", &choice); err != nil {
			if err == io.EOF {
				return ""
			}
			log.Fatalln("Error reading option:", err)
		}
		fmt.Println(choice, chosenArc.Options[0].Arc)
		if nextArc := chosenArc.Options[choice-1].Arc; nextArc != "" {
			return nextArc
		}
	}
	return ""
}

func main() {
	storyMap, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		log.Fatalf("Unable to read file: %s", err)
	}
	story := parseStory(storyMap)
	for nextArc := traverseArcs(story, "intro"); nextArc != ""; nextArc = traverseArcs(story, nextArc) {
	}
	fmt.Fprintf(os.Stdout, "%s\n", "End of adventure.")
}
