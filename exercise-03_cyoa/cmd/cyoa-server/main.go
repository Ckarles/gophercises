package main

import (
	"cyoa"
	"log"
	"os"
)

func main() {
	var story cyoa.Story

	f, err := os.Open("./gopher.json")
	defer f.Close()
	if err != nil {
		log.Println(err)
		return
	}

	err = story.FromJSON(f)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%+v", story)
}
