package main

import (
	"fmt"

	"github.com/fenriz07/cyoa/helpers"
	"github.com/fenriz07/cyoa/models/book"
)

func main() {
	fmt.Printf("%v \n", "Iniciando ejercicio #3")

	b := book.Init("book.json")

	cap := b.GetCap("intro")

	title := cap.GetTitle()
	//story := cap.GetStory()
	options := cap.GetOptions()

	helpers.DD(options[0].Arc)
	helpers.DD(title)
}
