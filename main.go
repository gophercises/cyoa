package main

import (
	"fmt"
	"net/http"

	"github.com/fenriz07/cyoa/handler"
)

func main() {
	fmt.Printf("%v \n", "Listen Serve 8080 := ")

	mux := defaultMux()

	nmux := handler.Init(mux)

	//b := book.Init("book.json")

	//cap := b.GetCap("intro")
	http.ListenAndServe(":8080", nmux)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Hello, world!")
}
