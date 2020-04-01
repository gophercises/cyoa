package httpstoryhandler

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"unicode/utf8"
)

// Handler used to serve the HTML response
type Handler struct {
	Scenarios map[string]Scenario
}

type Scenario struct {
	Title   string   `json:"title"`
	Story   []string `json: "story"`
	Options []option `json:"options"`
}

type option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	sceneTitle := sanitiseURLPath(req.URL.Path)
	scene, errFindingScene := retrieveScenarioFromMapOfScenarios(sceneTitle, h.Scenarios)
	var textResponse string
	if errFindingScene == nil {
		textResponse = createHTMLResponseForScenario(scene)
	} else {
		textResponse = "<p>You are doing really well.</p>"
	}
	writeTextToHTTPResponse(textResponse, w)
}

func sanitiseURLPath(urlPath string) string {
	// Path always contains a leading / , even when request is made without any
	var sanitised = trimFirstRune(urlPath)
	sanitised = strings.ToLower(sanitised)
	return sanitised
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func writeTextToHTTPResponse(text string, w http.ResponseWriter) {
	response := []byte(text)
	w.Write(response)
}

func createHTMLResponseForScenario(scene Scenario) string {
	tmpl, err := template.ParseFiles("./base.html")
	if err != nil {
		panic(err)
	}
	var textTemplate bytes.Buffer
	tmpl.Execute(&textTemplate, scene)
	return textTemplate.String()
}

func retrieveScenarioFromMapOfScenarios(scenarioTitle string, mapOfScenarios map[string]Scenario) (Scenario, error) {
	var err error
	scene, foundScenario := mapOfScenarios[scenarioTitle]
	if !foundScenario {
		err = fmt.Errorf("Unable to find a scenario of title \"%s\"", scenarioTitle)
	}
	return scene, err
}
