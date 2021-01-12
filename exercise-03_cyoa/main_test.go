package main

import (
	"strings"
	"testing"
)

func TestStoryLoadJSON(t *testing.T) {
	t.Run("story=safe", func(t *testing.T) {
		json := `
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

		st := Story{}
		st.LoadJSON(strings.NewReader(json))

		if in, ok := st["intro"]; !ok {
			t.Errorf("st[\"intro\"] missing")
		} else if got, want := in.Title, "The Little Blue Gopher"; got != want {
			t.Errorf("st[\"intro\"].title = %s; want: %s", got, want)
		}
	})

	t.Run("story=empty", func(t *testing.T) {
		json := `[]`

		st := Story{}
		err := st.LoadJSON(strings.NewReader(json))
		if err == nil {
			t.Error("Empty JSON was parsed as valid Story")
		}
	})
}
