package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// mockLoadJSON returns a mock of os.Open() for unit testing
func (st *Story) mockLoadJSON(js string) error {
	return st.LoadJSON("", func(path string) (io.ReadCloser, error) {
		return ioutil.NopCloser(strings.NewReader(js)), nil
	})
}

func TestStoryLoadJSON(t *testing.T) {

	safeJSON := `
{
"intro": {
	"title": "The Little Blue Gopher",
	"story": [
		"Visit..."
	],
	"options": [
		{
			"text": "New York",
			"arc": "new-york"
		},
		{
			"text": "Denver",
			"arc": "denver"
		}
	]
}
}
`

	t.Run("file=exists", func(t *testing.T) {
		if testing.Short() {
			t.Skip("skipping test in short mode.")
		}

		// Create temp file to host JSON
		tmpDirname, err := ioutil.TempDir("", "cyoa-")
		if err != nil {
			log.Fatalf("Cannot create temp dir: %s\n", err)
		}

		// remove tmpDir at the end of the test(s)
		defer func() {
			err := os.RemoveAll(tmpDirname)
			if err != nil {
				log.Fatalf("Cannot remove temp dir: %s\n", err)
			}
		}()

		// write safe JSON data
		filename := filepath.Join(tmpDirname, "safe-story.json")
		if err := ioutil.WriteFile(filename, []byte(safeJSON), 0666); err != nil {
			log.Fatalf("Cannot write temp file: %s", err)
		}

		// Set storyFile Flag manually
		*storyFile = filename
		main()
	})
	t.Run("file=missing", func(t *testing.T) {
		if testing.Short() {
			t.Skip("skipping test in short mode.")
		}

		// Set storyFile Flag manually
		*storyFile = ""
		main()
	})

	t.Run("story=safe", func(t *testing.T) {
		st := Story{}
		st.mockLoadJSON(safeJSON)

		if got, want := st["intro"].Title, "The Little Blue Gopher"; got != want {
			t.Errorf("st[\"intro\"].title = %s; want: %s", got, want)
		}
	})

	t.Run("story=empty", func(t *testing.T) {
		st := Story{}
		err := st.mockLoadJSON("[]")

		if err == nil {
			t.Error("Empty JSON was parsed as valid Story")
		}
	})
}
