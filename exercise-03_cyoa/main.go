package main

import (
	"encoding/json"
	"io"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []struct {
		Text    string `json:"test"`
		Chapter string `json:"arc"`
	} `json:"options"`
}

// LoadJSON loads a Story from a JSON File
func (story *Story) LoadJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(story)
}
