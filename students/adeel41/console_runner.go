package main

import (
	"fmt"
	"log"
	"os"
)

//ConsoleRunner is a type of `Runner` which displays story in command prompt
type ConsoleRunner struct {
}

//Start starts the runner by displaying the first story endpoint and then carry on from there
func (cr ConsoleRunner) Start(provider *StoryArcProvider) {
	cr.displayArcText(*provider, "intro")
}

func (cr ConsoleRunner) displayArcText(provider StoryArcProvider, arcName string) {

	arc, err := provider.WriteTemplatedText(os.Stdout, arcName)
	if err != nil {
		log.Println(err)
	}
	if len(arc.Options) == 0 {
		return
	}
	fmt.Print("Your Option: ")
	var optionNumber int
	fmt.Scan(&optionNumber)
	for _, option := range arc.Options {
		if option.Number == optionNumber {
			cr.displayArcText(provider, option.Arc)
		}
	}
}
