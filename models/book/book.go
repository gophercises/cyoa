package book

import (
	"encoding/json"
	"io/ioutil"

	"github.com/fenriz07/cyoa/helpers"
	"github.com/fenriz07/cyoa/models/cap"
)

type Book struct {
	Caps map[string]cap.Cap
}

func (b Book) GetCap(namecap string) cap.Cap {
	cap := b.Caps[namecap]

	return cap
}

func Init(namefile string) Book {

	bytejsonbook, err := ioutil.ReadFile(namefile)

	if err != nil {
		helpers.Exit(err)
	}

	caps := map[string]cap.Cap{}

	err = json.Unmarshal(bytejsonbook, &caps)

	if err != nil {
		helpers.Exit(err)
	}

	b := Book{Caps: caps}

	return b

}
