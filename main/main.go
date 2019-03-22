package main

import (
	"flag"
	"gopherex/cyoa/main/route-handler"
	"gopherex/cyoa/main/settings"
)

const (
	storyTemplate = "main/templates/story.html"
)

var (
	filePath = flag.String("filePath","","path to the story source file")
	listenPort = flag.String("listenPort","","port number that app will listen")
)

func main()  {
	flag.Parse()
	config := settings.Settings{FilePath: *filePath, ListenPort: *listenPort, TemplatePath: storyTemplate}
	handler := route_handler.RouteHandler{Settings: config}
	handler.ServeRequests()
}