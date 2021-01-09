package urlshort

import (
	"errors"
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToURLs map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if target, ok := pathsToURLs[r.URL.Path]; ok {
			http.Redirect(w, r, target, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathMap, err := pathMapFromYAML(yaml)
	if err != nil {
		return nil, err
	}
	return MapHandler(pathMap, fallback), nil
}

type pathYAML struct {
	Path string
	URL  string
}

func pathMapFromYAML(b []byte) (map[string]string, error) {
	var YAMLMap []pathYAML
	err := yaml.Unmarshal(b, &YAMLMap)
	if err != nil {
		return nil, err
	}

	pathMap := make(map[string]string, len(YAMLMap))
	for _, k := range YAMLMap {
		p := k.Path
		if p == "" {
			return nil, errors.New("Missing Path")
		}
		pathMap[k.Path] = k.URL
	}
	return pathMap, nil
}
