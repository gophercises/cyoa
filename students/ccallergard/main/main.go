package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gophercises/cyoa/students/ccallergard"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var (
	jsonf    = flag.String("f", "gopher.json", "Adventure JSON file")
	dir      = flag.String("d", "_html", "HTML output directory")
	port     = flag.Int("p", 8080, "HTTP port")
	template = flag.String("t", "scene_template.html", "HTML template")
)

func main() {
	flag.Parse()

	cyoajson, err := ioutil.ReadFile(*jsonf)
	if err != nil {
		log.Fatal(err)
	}

	var adv cyoa.Adventure
	err = json.Unmarshal(cyoajson, &adv)
	if err != nil {
		log.Fatal(err)
	}

	err = cyoa.Generate(adv, *dir, *template)
	if err != nil {
		log.Fatal(err)
	}

	cyoaMux := http.NewServeMux()
	cyoaMux.Handle("/", http.RedirectHandler("/cyoa/intro.html", http.StatusPermanentRedirect))
	cyoaMux.Handle("/cyoa/", http.StripPrefix("/cyoa/", http.FileServer(http.Dir(*dir))))
	fmt.Println("HTTP server listening on", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), cyoaMux))
}
