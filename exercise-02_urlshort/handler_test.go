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
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	// generate handler
	handler := MapHandler(pathsToUrls, defaultMux())

	t.Run("path=\"/yaml-godoc\"", func(t *testing.T) {
		rs, _ := sendReq(handler, "/yaml-godoc")

		t.Run("chk=statusCode", func(t *testing.T) {
			want := 301
			got := rs.StatusCode
			if got != want {
				t.Errorf("status code = \"%d\"; want \"%d\"", got, want)
			}
		})

		t.Run("chk=location", func(t *testing.T) {
			want := pathsToUrls["/yaml-godoc"]
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
