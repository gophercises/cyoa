package route_handler

import (
	"fmt"
	"gopherex/cyoa/students/cherednichenkoa/settings"
	"gopherex/cyoa/students/cherednichenkoa/source"
	"html/template"
	"net/http"
	"strings"
)

const (
	defaultUrl = "/intro"
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

		http.Redirect(w, req, rh.getRedirectUrl(req), http.StatusMovedPermanently)
	}
}

func (rh *RouteHandler) getPort() string {
	port := ":" + rh.Settings.GetListenPort()
	return port
}

func (rh *RouteHandler) getRedirectUrl(req *http.Request) string {
	redirectUrl := fmt.Sprintf("%s%s%s", "http://", req.Host, defaultUrl)
	return redirectUrl
}

func (rh *RouteHandler) prepareUrl(req *http.Request) string {
	path := strings.Trim(req.URL.Path,"/")
	return path
}