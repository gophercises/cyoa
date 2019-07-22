package main

import (
	"fmt"
	"net/http"
)

//WebRunner is a type of `Runner` which displays story on a web page
type WebRunner struct {
}

//Start starts the runner by displaying the first story endpoint and then carry on from there
func (wr WebRunner) Start(provider *StoryArcProvider) {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		wr.rootEndpointHandler(provider, w, req)
	})
	fmt.Println("Listening on 8888 port .....")
	http.ListenAndServe(":8888", nil)

}

func (wr WebRunner) rootEndpointHandler(provider *StoryArcProvider, w http.ResponseWriter, req *http.Request) {

	arcValues := req.URL.Query()["arc"]
	arcEndpoint := "intro"
	if len(arcValues) > 0 {
		arcEndpoint = arcValues[0]
	}

	_, err := provider.WriteTemplatedText(w, arcEndpoint)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
