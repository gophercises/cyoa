package main

//Runner is implemented by ConsoleRunner and WebRunner
type Runner interface {
	Start(provider *StoryArcProvider)
}
