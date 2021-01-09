package story

import (
	"errors"
	"testing"
)

// mock ioutil.ReadFile
func init() {
	readFile = func(path string) ([]byte, error) {
		switch path {
		case "safe":
			return []byte(`
				{
				  "intro": {
				    "title": "The Little Blue Gopher",
				    "story": [
				      "Once upon a time, long long ago, there was a little blue gopher. Our little blue friend wanted to go on an adventure, but he wasn't sure where to go. Will you go on an adventure with him?",
				      "One of his friends once recommended going to New York to make friends at this mysterious thing called \"GothamGo\". It is supposed to be a big event with free swag and if there is one thing gophers love it is free trinkets. Unfortunately, the gopher once heard a campfire story about some bad fellas named the Sticky Bandits who also live in New York. In the stories these guys would rob toy stores and terrorize young boys, and it sounded pretty scary.",
				      "On the other hand, he has always heard great things about Denver. Great ski slopes, a bad hockey team with cheap tickets, and he even heard they have a conference exclusively for gophers like himself. Maybe Denver would be a safer place to visit."
				    ],
				    "options": [
				      {
				        "text": "That story about the Sticky Bandits isn't real, it is from Home Alone 2! Let's head to New York.",
				        "arc": "new-york"
				      },
				      {
				        "text": "Gee, those bandits sound pretty real to me. Let's play it safe and try our luck in Denver.",
				        "arc": "denver"
				      }
				    ]
				  }
				}
			`), nil
		case "unsafe":
			return []byte(`[]`), nil
		default:
			return nil, errors.New("")
		}
	}
}

func TestStoryFromJSON(t *testing.T) {
	t.Run("file=readable", func(t *testing.T) {
		t.Run("json=safe", func(t *testing.T) {
			st := Story{}
			st.FromJSON("safe")

			if in, ok := st["intro"]; !ok {
				t.Errorf("st[\"intro\"] missing")

			} else if got, want := in.Title, "The Little Blue Gopher"; got != want {
				t.Errorf("st[\"intro\"].title = %s; want: %s", got, want)
			}
		})

		t.Run("json=unsafe", func(t *testing.T) {
			st := Story{}
			err := st.FromJSON("unsafe")

			if err == nil {
				t.Error("Malformed JSON error uncaught")
			}
		})
	})

	t.Run("file=unreadable", func(t *testing.T) {
		st := Story{}
		err := st.FromJSON("unreadable")

		if err == nil {
			t.Error("File reading error uncaught")
		}
	})

}
