package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Problem struct {
	question string
	answer   string
}

// getProblems returns a slice a problems from a csv file
func getProblems(path string) ([]Problem, error) {

	// open the file
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	// read the file as csv
	c := csv.NewReader(file)
	// init returned slice
	var ps []Problem
	for {
		// for each line
		r, err := c.Read()

		if err == io.EOF {
			// EOF reached, return the slice
			return ps, nil

		} else if err != nil {
			// any other error, return it
			return nil, err

		} else {
			// add new Problem to slice
			ps = append(ps, Problem{r[0], r[1]})
		}
	}
}

func main() {
	problems, err := getProblems("problems.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(problems)
}
