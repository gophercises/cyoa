package main

//ArcOption is an option for the story
type ArcOption struct {
	Number int
	Text   string
	Arc    string
}

//Load loads the json data into an ArcOption
func (ao *ArcOption) Load(number int, data map[string]interface{}) {
	ao.Number = number
	ao.Text = data["text"].(string)
	ao.Arc = data["arc"].(string)
}
