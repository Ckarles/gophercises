package urlshort

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func sendReq(handler http.HandlerFunc, target string) (*http.Response, string) {
	rq := httptest.NewRequest("GET", target, nil)
	w := httptest.NewRecorder()
	handler(w, rq)

	rs := w.Result()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}
	return rs, string(body)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
	return mux
}

func TestMapHandler(t *testing.T) {
	// mock redirect map
	pathsToURLs := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	// generate handler
	handler := MapHandler(pathsToURLs, defaultMux())

	t.Run("path=\"/yaml-godoc\"", func(t *testing.T) {
		rs, _ := sendReq(handler, "/yaml-godoc")

		t.Run("chk=statusCode", func(t *testing.T) {
			want := http.StatusFound
			got := rs.StatusCode
			if got != want {
				t.Errorf("status code = \"%d\"; want \"%d\"", got, want)
			}
		})

		t.Run("chk=location", func(t *testing.T) {
			want := pathsToURLs["/yaml-godoc"]
			got := rs.Header.Get("Location")
			if got != want {
				t.Errorf("Header[\"Location\"] = \"%s\"; want \"%s\"", got, want)
			}
		})
	})

	t.Run("path=\"/xyz\"", func(t *testing.T) {
		_, body := sendReq(handler, "/xyz")

		t.Run("chk=body", func(t *testing.T) {
			want := "Hello, world!"
			got := strings.TrimSpace(body)
			if want != got {
				t.Errorf("Response Body = \"%s\"; want \"%s\"", got, want)
			}
		})
	})
}

func TestYAMLHandler(t *testing.T) {

	t.Run("yaml=safe", func(t *testing.T) {
		yaml := []byte(`
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`)

		handler, err := YAMLHandler(yaml, defaultMux())
		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}

		rs, _ := sendReq(handler, "/urlshort")
		t.Run("chk=statusCode", func(t *testing.T) {
			want := http.StatusFound
			got := rs.StatusCode
			if got != want {
				t.Errorf("status code = \"%d\"; want \"%d\"", got, want)
			}
		})

		t.Run("chk=location", func(t *testing.T) {
			want := "https://github.com/gophercises/urlshort"
			got := rs.Header.Get("Location")
			if got != want {
				t.Errorf("Header[\"Location\"] = \"%s\"; want \"%s\"", got, want)
			}
		})
	})

	t.Run("yaml=malformed", func(t *testing.T) {
		yaml := []byte(`
url: /fefg
path: xyz
`)
		_, err := YAMLHandler(yaml, defaultMux())
		if err == nil {
			t.Errorf("Malformed yaml has been parsed (it shouldn't have)")
		}
	})

	t.Run("yaml=miss-path", func(t *testing.T) {
		yaml := []byte(`
- url: /fefg
`)
		_, err := YAMLHandler(yaml, defaultMux())
		if err == nil {
			t.Errorf("Malformed yaml has been parsed (it shouldn't have)")
		}
	})
}
