package story

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

func (story *Story) FromJSON(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&story)
	if err != nil {
		return err
	}
	return nil
}
