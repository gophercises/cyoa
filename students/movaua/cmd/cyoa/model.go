package main

// Book is Choose Your Own Adventure book
type Book map[string]Chaper

// Chaper of a book
type Chaper struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

// Option of a Chaper
type Option struct {
	Text   string `json:"text"`
	Chaper string `json:"arc"`
}
