package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type customHandler struct {
}

type option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type story struct {
	Title   string   `json:"title"`
	Story   []string `json: "story"`
	Options []option `json:"options"`
}

func (h customHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	textResponse := "<p>You are doing really well.</p>"
	response := []byte(textResponse)
	w.Write(response)
}

func main() {
	mux := http.NewServeMux()
	defaultHandler := customHandler{}
	mux.Handle("/", defaultHandler)

	jsonStories, err := ioutil.ReadFile("./stories.json")
	if err != nil {
		panic(err)
	}
	var stories map[string]story
	unmarshallingError := json.Unmarshal(jsonStories, &stories)
	if unmarshallingError != nil {
		panic(unmarshallingError)
	}
	fmt.Println(stories)

	fmt.Println("Launching server on port 3645")
	err = http.ListenAndServe(":3645", mux)
	if err != nil {
		panic(err)
	}
}
