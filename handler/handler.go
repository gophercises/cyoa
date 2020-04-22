package handler

import (
	"html/template"
	"net/http"

	"github.com/fenriz07/cyoa/helpers"
	"github.com/fenriz07/cyoa/models/book"
)

func Init(fallback http.Handler) http.HandlerFunc {

	views := template.Must(template.ParseGlob("resources/views/*"))

	b := book.Init("book.json")

	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {

		key, ok := request.URL.Query()["slug"]

		var slug string

		if !ok || len(key[0]) < 1 {
			slug = "intro"
		} else {
			slug = key[0]
		}

		cap := b.GetCap(slug)

		err := views.ExecuteTemplate(w, "story.html", cap)

		if err != nil {
			helpers.DD(err)
		}

		fallback.ServeHTTP(w, request)
	})
}
