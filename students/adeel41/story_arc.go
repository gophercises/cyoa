package main

import (
	"bytes"
)

//StoryArc is an arc of the story which also contains the options to go to other arcs
type StoryArc struct {
	Identifier string
	Title      string
	Paragraph  string
	Options    []*ArcOption
}

//Load loads the json data and set StoryArc
func (sa *StoryArc) Load(key string, data map[string]interface{}) {
	var buffer bytes.Buffer
	paragraphs := data["story"].([]interface{})
	for _, p := range paragraphs {
		buffer.WriteString(p.(string) + "\r\n")
	}

	sa.Identifier = key
	sa.Title = data["title"].(string)
	sa.Paragraph = buffer.String()

	var options []*ArcOption
	for i, v := range data["options"].([]interface{}) {
		opt := new(ArcOption)
		opt.Load(i+1, v.(map[string]interface{}))
		options = append(options, opt)
	}
	sa.Options = options
}
