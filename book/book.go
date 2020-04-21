package book

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Book struct {
}

type Cap struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func ReadBook() {

	bytejsonbook, err := ioutil.ReadFile("book.json")

	if err != nil {
		ddExit(err)
	}

	m := map[string]Cap{}

	err = json.Unmarshal(bytejsonbook, &m)

	if err != nil {
		ddExit(err)
	}

	dd(m["intro"])

}

func dd(any interface{}) {
	fmt.Printf("El valor de la variable es %v \n tipo %T\n", any, any)
	os.Exit(2)
}

func ddExit(e error) {
	fmt.Println(e)
	os.Exit(2)
}
