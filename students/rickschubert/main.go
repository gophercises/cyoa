package main

import (
	"fmt"
	"net/http"
)

type customHandler struct {
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

	fmt.Println("Launching server on port 3645")
	err := http.ListenAndServe(":3645", mux)
	if err != nil {
		panic(err)
	}
}
