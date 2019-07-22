package cyoa

import (
	"strings"
	"testing"
)

func TestParseJSON(t *testing.T) {
	story := `{
		"intro": {
		  "title": "The Little Blue Gopher",
		  "story": [
			"Once upon a time,?",
			"On the other h"
		  ],
		  "options": [
			{
			  "text": "That story about the ",
			  "arc": "new-york"
			},
			{
			  "text": "Gee, those b",
			  "arc": "home"
			}
		  ]
		},
		"new-york": {
		  "title": "Visiting New York",
		  "story": [
			"Upon arrivi d.",
			"As yto?"
		  ],
		  "options": [
			{
			  "text": "This is ge.",
			  "arc": "home"
			},
			{
			  "text": "Maybe people .",
			  "arc": "intro"
			}
		  ]
		},
	  
		"home": {
		  "title": "Home Sweet Home",
		  "story": [
			"Your little gopher bna."
		  ],
		  "options": []
		}
	  }`

	_, err := ParseJSON(strings.NewReader(story))
	if err != nil {
		t.Error(err)
	}

}
