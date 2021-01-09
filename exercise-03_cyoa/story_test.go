package cyoa

import (
	"strings"
	"testing"
)

func TestStoryFromJSON(t *testing.T) {
	t.Run("file=readable", func(t *testing.T) {
		t.Run("json=safe", func(t *testing.T) {
			st := Story{}
			r := strings.NewReader(`
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
`)
			st.FromJSON(r)

			if in, ok := st["intro"]; !ok {
				t.Errorf("st[\"intro\"] missing")

			} else if got, want := in.Title, "The Little Blue Gopher"; got != want {
				t.Errorf("st[\"intro\"].title = %s; want: %s", got, want)
			}
		})

		t.Run("json=unsafe", func(t *testing.T) {
			st := Story{}
			r := strings.NewReader("[]")
			err := st.FromJSON(r)

			if err == nil {
				t.Error("Malformed JSON error uncaught")
			}
		})
	})
}
