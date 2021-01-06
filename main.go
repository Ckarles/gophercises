package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Problem struct {
	question string
	answer   string
}

// getProblems returns a slice a problems from a csv file
func getProblems(path string) ([]Problem, error) {

	var file *os.File
	var err error

	// else open file from path
	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

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
			ps = append(ps, Problem{
				question: r[0],
				answer:   strings.TrimSpace(r[1]),
			})
		}
	}
}

// declare cli flags
var pathFlag = flag.String("csv", "problems.csv", "a csv file in the format question,answer")

func main() {

	// parse cli flags
	flag.Parse()

	problems, err := getProblems(*pathFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
