package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

const FirstChapterTag string = "intro"

// A Continuation contains data about the chapter that might follow the current
// chapter, including its tag and a preview of its text.
type Continuation struct {
	Preview   string `json:"text"`
	TargetTag string `json:"arc"`
}

type Chapter struct {
	Title         string         `json:"title"`
	Text          []string       `json:"story"`
	Continuations []Continuation `json:"options"`
}

type Book map[string]Chapter

func (book Book) getChapter(chapterTag string) (Chapter, error) {
	chapter, ok := book[chapterTag]
	if ok {
		return chapter, nil
	}

	return chapter, getNoChapterError(chapterTag)
}

func getNoChapterError(chapterTag string) error {
	msg := fmt.Sprintf("Chapter with tag %s does not exist.", chapterTag)
	return errors.New(msg)
}

func fromJSON(data []byte) Book {
	var book Book
	json.Unmarshal(data, &book)

	return book
}
