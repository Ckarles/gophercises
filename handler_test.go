package urlshort

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMapHandler(t *testing.T) {
	// test handler
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})

	handler := MapHandler(pathsToUrls, mux)

	// test defined route
	rq := httptest.NewRequest("GET", "/yaml-godoc", nil)
	w := httptest.NewRecorder()
	handler(w, rq)

	rs := w.Result()

	// check status code
	sc_want := 301
	sc_got := rs.StatusCode
	if sc_got != sc_want {
		t.Errorf("status code = \"%d\"; want \"%d\"", sc_got, sc_want)
	}

	// check Location
	l_want := pathsToUrls["/yaml-godoc"]
	l_got := rs.Header.Get("Location")
	if l_got != l_want {
		t.Errorf("Header[\"Location\"] = \"%s\"; want \"%s\"", l_got, l_want)
	}

	// test fallback route
	rq = httptest.NewRequest("GET", "/test", nil)
	w = httptest.NewRecorder()
	handler(w, rq)

	rs = w.Result()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	// check body
	b_want := "Hello, world!"
	b_got := strings.TrimSpace(string(body))
	if b_want != b_got {
		t.Errorf("Response Body = \"%s\"; want \"%s\"", b_got, b_want)
	}
}
