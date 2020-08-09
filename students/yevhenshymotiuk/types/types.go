// Package types provides custom types
package types

// Option provides data for user selection
type Option struct {
	Text string
	Arc  string
}

// Page describes entire adventure page
type Page struct {
	Title   string
	Story   []string
	Options []Option
}
