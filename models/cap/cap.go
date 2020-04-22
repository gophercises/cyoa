package cap

type Cap struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func (c Cap) GetTitle() string {
	return c.Title
}

func (c Cap) GetStory() []string {
	return c.Story
}

func (c Cap) GetOptions() []struct {
	Text string "json:\"text\""
	Arc  string "json:\"arc\""
} {

	return c.Options
}
