package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

type server struct {
	port   string
	router *mux.Router
}

func newServer() *server {
	router := mux.NewRouter()
	return &server{
		router: router,
		port:   ":3000",
	}
}

func (s *server) Index(w http.ResponseWriter, r *http.Request) {
	Stories, err := getStories()
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("webpage").Parse(IndexHtml)
	t.Execute(w, map[string]any{
		"Stories": Stories,
	})
}

func (s *server) Story(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if title, ok := vars["title"]; ok {
		Story, err := getStorieByTitle(title)
		if err != nil {
			log.Println(err)
			w.Write([]byte(err.Error()))
			return
		}
		t, err := template.New("webpage").Parse(StoryHtml)
		t.Execute(w, map[string]any{
			"ArcTitle": title,
			"Story":    Story,
		})
	} else {
		w.Write([]byte("..."))
	}
}

func (s *server) setupRoutes() {
	s.router.HandleFunc("/stories", s.Index)
	s.router.HandleFunc("/story/{title}", s.Story)
}

func (s *server) start() {
	http.ListenAndServe(""+s.port, s.router)
}
