package route_handler

import (
	"gopherex/cyoa/students/cherednichenkoa/settings"
	"gopherex/cyoa/students/cherednichenkoa/source"
	"html/template"
	"net/http"
	"strings"
)

const (
	defaultStory = "intro"
)

type RouteHandler struct {
	Settings settings.Settings
}

func (rh *RouteHandler) ServeRequests() {
	fileHandler := source.JsonFileHandler{Settings: rh.Settings}
	fileContent, err := fileHandler.GetFileContent()
	if err != nil {
		panic(err)
	}
	urlHandler := rh.getMapHandler(fileContent)
	http.HandleFunc("/", urlHandler)
	http.ListenAndServe(rh.getPort(), nil)
}

func (rh *RouteHandler) getMapHandler(stories map[string]source.StoryDetails) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		tmpl := template.Must(template.ParseFiles(rh.Settings.GetTemplatePath()))
		path := rh.prepareUrl(req)
		story, ok := stories[path]
		if ok {
			tmpl.Execute(w, story)
			return
		}

		tmpl.Execute(w, stories[defaultStory])
	}
}

func (rh *RouteHandler) getPort() string {
	port := ":" + rh.Settings.GetListenPort()
	return port
}

func (rh *RouteHandler) prepareUrl(req *http.Request) string {
	path := strings.Trim(req.URL.Path,"/")
	return path
}