package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var cyoaTemplate = makeTemplate("cyoa.html")

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

func cliHandler(s StoryArc) string {
	fmt.Fprintln(os.Stdout, s.explainStoryArc())
	if len(s.Options) > 0 {
		fmt.Fprintln(os.Stdout, "Choose an option:")
		for index, option := range s.Options {
			fmt.Fprintf(os.Stdout, "\n%d %s: %s\n", index+1, option.Arc, option.Text)
		}
		var choice int
		if _, err := fmt.Fscanf(os.Stdin, "%d", &choice); err != nil {
			if err == io.EOF {
				return ""
			}
			log.Fatalln("Error reading option:", err)
		}
		fmt.Fprintln(os.Stdout, choice, s.Options[0].Arc)
		if nextArc := s.Options[choice-1].Arc; nextArc != "" {
			return nextArc
		}
	}
	return ""
}

func makeTemplate(filename string) *template.Template {
	return template.Must(template.ParseFiles(filename))
}

func (story Story) httpHandler(arc string, w http.ResponseWriter) error {
	t, err := cyoaTemplate.Clone()
	if err != nil {
		return err
	}
	data := story[arc]
	// log.Printf("%+v", data)
	return t.Execute(w, data)
}

// Story is an adventure with story arcs
type Story map[string]StoryArc

func (story Story) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arc := strings.TrimSpace(r.URL.Path)
	if arc == "/" {
		arc = "/intro"
	}
	if strings.HasSuffix(arc, "/") {
		arc = strings.TrimSuffix(arc, "/")
	}
	if err := story.httpHandler(arc[1:], w); err != nil {
		log.Println("error rendering template", err)
	}
}

func parseStory(storyMap []byte) (story Story) {
	if err := json.Unmarshal(storyMap, &story); err != nil {
		log.Fatalf("Error parsing story map: %s", err)
	}
	return
}

func (story Story) traverseArcs(arc string, action func(StoryArc) string) {
	chosenArc := story[arc]
	if nextArc := action(chosenArc); nextArc != "" {
		story.traverseArcs(nextArc, action)
	}
}

func main() {
	storyMap, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		log.Fatalf("Unable to read file: %s", err)
	}
	story := parseStory(storyMap)
	// story.traverseArcs("intro", cliHandler)
	// fmt.Fprintf(os.Stdout, "%s\n", "End of adventure.")

	r := http.DefaultServeMux

	r.Handle("/", story)

	fmt.Fprintln(os.Stderr, http.ListenAndServe(":3000", story))
}
