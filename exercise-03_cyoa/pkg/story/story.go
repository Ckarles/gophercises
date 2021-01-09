package story

import (
	"encoding/json"
	"io/ioutil"
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

var readFile = ioutil.ReadFile

func (story *Story) FromJSON(path string) error {
	data, err := readFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &story)
	if err != nil {
		return err
	}

	return nil
}
