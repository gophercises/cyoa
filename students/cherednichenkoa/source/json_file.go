package source

import (
	"encoding/json"
	"fmt"
	"gopherex/cyoa/students/cherednichenkoa/settings"
	"io/ioutil"
)

type JsonFileHandler struct  {
	Settings settings.Settings
}

func (fh *JsonFileHandler) GetFileContent () (map[string]StoryDetails, error) {
	file, err := ioutil.ReadFile(fh.Settings.GetFilePath())
	if err != nil {
		fmt.Println("Error during json file reading.")
		panic(err)
	}
	var out map[string]StoryDetails
	if err := json.Unmarshal(file, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type StoryOption struct {
	Text      string `json:"text"`
	Arc string `json:"arc"`
}

type StoryDetails struct {
	Title string `json:"title"`
	Story []string  `json:"story"`
	Options []StoryOption `json:"options"`
}
