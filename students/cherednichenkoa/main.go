package main

import (
	"flag"
	"gopherex/cyoa/students/cherednichenkoa/route-handler"
	"gopherex/cyoa/students/cherednichenkoa/settings"
)

const (
	storyTemplate = "cherednichenkoa/templates/story.html"
)

var (
	filePath = flag.String("filePath","","path to the story source file")
	listenPort = flag.String("listenPort","","port number that app will listen")
)

func main()  {
	flag.Parse()
	if len(*listenPort) == 0 || len(*filePath) == 0 {
		panic("Please specify application params (listenPort and filePath)")
	}
	config := settings.Settings{FilePath: *filePath, ListenPort: *listenPort, TemplatePath: storyTemplate}
	handler := route_handler.RouteHandler{Settings: config}
	handler.ServeRequests()
}