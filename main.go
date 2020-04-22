package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("%v \n", "Iniciando ejercicio #3")

	//b := book.Init("book.json")

	//cap := b.GetCap("intro")
	http.ListenAndServe(":8080", defaultMux())

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
