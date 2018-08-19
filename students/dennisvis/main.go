package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/alecthomas/template"
)

var useCLI = flag.Bool("useCLI", false, "Wheher or not to use the CLI version of the game, web-based otherwise")

type storyOption struct {
	Text string
	Arc  string
}

type storyArc struct {
	Title   string
	Story   []string
	Options []storyOption
}

type storyHandler struct {
	storyArcs map[string]storyArc
	template  *template.Template
}

func (sh storyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var path string
	if req.URL.Path == "/" {
		path = "intro"
	} else {
		path = strings.TrimLeft(req.URL.Path, "/")
	}
	storyArc := sh.storyArcs[path]
	sh.template.Execute(res, storyArc)
}

type storyCLI struct {
	storyArcs map[string]storyArc
	reader    *bufio.Reader
}

func (scli storyCLI) getStoryOption(options []storyOption) storyOption {
	nrOfChoices := len(options)

	choice, err := scli.reader.ReadString('\n')
	if err != nil {
		fmt.Println("Sorry, you're choice could not be read, please try again...")
		return scli.getStoryOption(options)
	}

	choiceNr, err := strconv.Atoi(strings.TrimRight(choice, "\n"))
	if err != nil {
		fmt.Println("Sorry, you're choice has to be number, please try again...")
		return scli.getStoryOption(options)
	}
	if choiceNr <= 0 || choiceNr > nrOfChoices {
		fmt.Printf("Sorry, you're choice has to be number between %d and %d, please try again...\n", 1, nrOfChoices)
		return scli.getStoryOption(options)
	}

	return options[choiceNr-1]
}

func (scli storyCLI) presentStoryArc(storyArcName string) {
	storyArc := scli.storyArcs[storyArcName]

	fmt.Printf("\n--- %s ---\n\n", storyArc.Title)

	for _, p := range storyArc.Story {
		fmt.Println(p)
	}

	if len(storyArc.Options) == 0 {
		fmt.Printf("\n\nYour adventure has ended\n\n")
		os.Exit(0)
	}

	fmt.Println("\nWhat will you do?")
	for i, opt := range storyArc.Options {
		fmt.Printf("%d: %s\n", i+1, opt.Text)
	}
	fmt.Println("")

	so := scli.getStoryOption(storyArc.Options)
	scli.presentStoryArc(so.Arc)
}

func (scli storyCLI) start() {
	scli.presentStoryArc("intro")
}

func main() {
	flag.Parse()

	f, err := os.Open("gopher.json")
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(f)
	if err != nil {
		panic(err)
	}

	var storyArcs map[string]storyArc
	err = json.Unmarshal(buf.Bytes(), &storyArcs)
	if err != nil {
		panic(err)
	}

	if *useCLI {
		storyCLI{storyArcs, bufio.NewReader(os.Stdin)}.start()
	} else {
		t, err := template.ParseFiles("templates/main.html")
		if err != nil {
			panic(err)
		}

		fmt.Println("Your adventure awaits, open your browser and visit localhost:8080 to begin")
		http.ListenAndServe(":8080", storyHandler{storyArcs, t})
	}
}
