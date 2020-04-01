package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophersizes/cyoa/students/rickschubert/httpstoryhandler"
	"io/ioutil"
	"net/http"
	"strings"
)

func scenarioKeyNeedsToBeLowerCased(originalKey string, lowerCasedKey string) bool {
	return lowerCasedKey != originalKey
}

func lowerCaseScenarioKeys(m map[string]httpstoryhandler.Scenario) {
	for key, value := range m {
		lowerCasedKey := strings.ToLower(key)
		if scenarioKeyNeedsToBeLowerCased(key, lowerCasedKey) {
			m[strings.ToLower(key)] = value
			delete(m, key)
		}
	}
}

func parseScenariosFromFile() map[string]httpstoryhandler.Scenario {
	jsonStories, err := ioutil.ReadFile("./stories.json")
	if err != nil {
		panic(err)
	}
	var scenarios map[string]httpstoryhandler.Scenario
	unmarshallingError := json.Unmarshal(jsonStories, &scenarios)
	if unmarshallingError != nil {
		panic(unmarshallingError)
	}
	lowerCaseScenarioKeys(scenarios)
	return scenarios
}

func main() {
	scenarios := parseScenariosFromFile()

	mux := http.NewServeMux()
	defaultHandler := httpstoryhandler.Handler{
		Scenarios: scenarios,
	}
	mux.Handle("/", defaultHandler)

	fmt.Println("Launching server on port 3645")
	err := http.ListenAndServe(":3645", mux)
	if err != nil {
		panic(err)
	}
}
