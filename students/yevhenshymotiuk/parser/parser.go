// Package parser provides functions to get data from a file
package parser

import (
	"bufio"
	"encoding/json"
	"os"

	"cyoa/types"
)

// GetPages parses JSON to get map of pages
func GetPages(filePath string) (map[string]types.Page, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	r := bufio.NewReader(f)

	var objMap map[string]json.RawMessage

	decoder := json.NewDecoder(r)
	err = decoder.Decode(&objMap)
	if err != nil {
		return nil, err
	}

	pages := map[string]types.Page{}

	for place, pageJSON := range objMap {
		var page types.Page

		err := json.Unmarshal(pageJSON, &page)
		if err != nil {
			return nil, err
		}

		pages[place] = page
	}

	return pages, nil
}
