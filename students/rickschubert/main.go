package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type customHandler struct {
	Scenarios map[string]scenario
}

type option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type scenario struct {
	Title   string   `json:"title"`
	Story   []string `json: "story"`
	Options []option `json:"options"`
}

func createHtmlResponseForScenario(scene scenario) string {
	return scene.Title
}

func retrieveScenarioFromMapOfScenarios(scenarioTitle string, mapOfScenarios map[string]scenario) (scenario, error) {
	var err error
	scene, foundScenario := mapOfScenarios[scenarioTitle]
	if !foundScenario {
		err = errors.New(fmt.Sprintf("Unable to find a scenario of title \"%s\"", scenarioTitle))
	}
	return scene, err
}

func (h customHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	var textResponse string
	if req.URL.Path == "/intro" {
		// TODO: This assumes "intro" exists - needs to check whether it actually exists or not
		scene, _ := retrieveScenarioFromMapOfScenarios("intro", h.Scenarios)
		textResponse = createHtmlResponseForScenario(scene)
	} else {
		textResponse = "<p>You are doing really well.</p>"
	}
	response := []byte(textResponse)
	w.Write(response)
}

func main() {
	jsonStories, err := ioutil.ReadFile("./stories.json")
	if err != nil {
		panic(err)
	}
	var scenarios map[string]scenario
	unmarshallingError := json.Unmarshal(jsonStories, &scenarios)
	if unmarshallingError != nil {
		panic(unmarshallingError)
	}
	fmt.Println(scenarios)

	mux := http.NewServeMux()
	defaultHandler := customHandler{
		Scenarios: scenarios,
	}
	mux.Handle("/", defaultHandler)

	fmt.Println("Launching server on port 3645")
	err = http.ListenAndServe(":3645", mux)
	if err != nil {
		panic(err)
	}
}
