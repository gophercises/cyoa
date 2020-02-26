package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	parseInputs()

	launchBookServer(fromJSON(readFile(bookFilepath)))
}

var (
	bookFilepath         string
	htmlTemplateFilepath string
)

func parseInputs() {
	const defaultBookFilepath = "gopher.json"
	const defaultHTMLTemplateFilepath = "template.html"

	bkFilepath := flag.String("book", defaultBookFilepath, "filepath of the book")
	htmlTmplFilepath := flag.String("template", defaultHTMLTemplateFilepath, "filepath of the HTML template")

	flag.Parse()

	bookFilepath = *bkFilepath
	htmlTemplateFilepath = *htmlTmplFilepath
}

func getHTMLTemplate(templateFilepath string) *template.Template {
	tmpl, err := template.ParseFiles(templateFilepath)
	if err != nil {
		panic(err)
	}

	return tmpl
}

func readFile(fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func launchBookServer(book Book) {
	handler := newHandler(book)

	const serverPort = ":8080"
	err := http.ListenAndServe(serverPort, handler)
	if err != nil {
		panic(err)
	}
}

func newHandler(book Book) handlerHTTPBookServer {
	return handlerHTTPBookServer{book, getHTMLTemplate(htmlTemplateFilepath)}
}

type handlerHTTPBookServer struct {
	book         Book
	htmlTemplate *template.Template
}

func (h handlerHTTPBookServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	nextChapter := h.getNextChapter(r)
	h.htmlTemplate.Execute(w, nextChapter)
}

func (h handlerHTTPBookServer) getNextChapter(r *http.Request) Chapter {
	nextChapterTag := h.getNextChapterTag(r)

	chapter, err := h.book.getChapter(nextChapterTag)
	if err != nil {
		chapter, _ := h.book.getChapter(FirstChapterTag)
		return chapter
	}

	return chapter
}

func (h handlerHTTPBookServer) getNextChapterTag(r *http.Request) string {
	uri := r.RequestURI
	if len(uri) > 1 {
		// We strip off the leading forward slash from the URI.
		return r.RequestURI[1:]
	}

	return uri
}
