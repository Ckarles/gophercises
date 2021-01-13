package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
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

type fileOpener func(path string) (io.ReadCloser, error)

// LoadJSON loads a Story from a JSON File
func (story *Story) LoadJSON(path string, open fileOpener) error {
	// open json file
	r, err := open(path)
	if err != nil {
		return fmt.Errorf("Error opening story file: %v", err)
	}
	defer r.Close()

	// decode content
	if err := json.NewDecoder(r).Decode(story); err != nil {
		return fmt.Errorf("Error decoding story file: %v", err)
	}

	return nil
}

var storyFile = flag.String("story", "gopher.json", "JSON file containing the story")

func main() {
	st := Story{}
	st.LoadJSON("test", func(path string) (io.ReadCloser, error) {
		return os.Open(path)
	})
}
