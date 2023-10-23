package main

func main() {
	s := newServer()
	s.setupRoutes()
	s.start()
}
