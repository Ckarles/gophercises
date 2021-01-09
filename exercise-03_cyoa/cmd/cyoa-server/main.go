package main

import (
	"cyoa/pkg/story"
	"log"
	"os"
)

func main() {
	st := story.Story{}

	f, err := os.Open("./gopher.json")
	defer f.Close()
	if err != nil {
		log.Println(err)
		return
	}

	err = st.FromJSON(f)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%+v", st)
}
