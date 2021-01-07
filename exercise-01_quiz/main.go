package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Quizz []Problem

type Problem struct {
	question string
	answer   string
}

// getProblems returns a slice a problem from a csv file
func getQuizz(path string) (Quizz, error) {

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
	var qz Quizz
	for {
		// for each line
		r, err := c.Read()

		if err == io.EOF {
			// EOF reached, return the slice
			return qz, nil

		} else if err != nil {
			// any other error, return it
			return nil, err

		} else {
			// add new Problem to slice
			qz = append(qz, Problem{
				question: r[0],
				answer:   strings.TrimSpace(r[1]),
			})
		}
	}
}

// StartQuizz ask the questions and returns the score
func (quizz Quizz) Start() int {

	timer := time.NewTimer(time.Duration(*limitFlag) * time.Second)

	score := 0
	for i, p := range quizz {
		answerCh := make(chan string)
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		go func() {
			var a string
			fmt.Scanln(&a)
			answerCh <- a
		}()

		select {
		case a := <-answerCh:
			if a == p.answer {
				score++
			}
		case <-timer.C:
			fmt.Println()
			return score
		}
	}
	return score
}

// declare cli flags
var pathFlag = flag.String("csv", "problems.csv", "a csv file in the format question,answer")
var limitFlag = flag.Int("limit", 30, "the time limit for the quizz in seconds")

func main() {

	// parse cli flags
	flag.Parse()

	quizz, err := getQuizz(*pathFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	score := quizz.Start()
	fmt.Printf("You score %d out of %d.\n", score, len(quizz))
}
