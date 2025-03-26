package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func readUserInput(answerCh chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		answerCh <- scanner.Text()
	}
	close(answerCh)
}

func askQuestion(question string, timeLimit time.Duration, answerCh chan string) (string, bool) {
	fmt.Println(question)

	timer := time.NewTimer(timeLimit)
	defer timer.Stop()

	select {
	case <-answerCh:
	default:
	}

	select {
	case answer := <-answerCh:
		return answer, true
	case <-timer.C:
		fmt.Println("Time's up!")
		return "", false
	}
}

func main() {
	var filename string
	var limit int

	flag.StringVar(&filename, "csv", "problems.csv", "a csv file in the format of 'question, answer' (default 'problems.csv')")
	flag.IntVar(&limit, "limit", 30, "the time limit for quiz in seconds (default 30)")
	flag.Parse()

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	N := len(data)
	correctCnt := 0

	timeLimit := time.Duration(limit) * time.Second
	answerCh := make(chan string, 1)

	go readUserInput(answerCh)

	for i, row := range data {
		ques := fmt.Sprintf("Problem #%d: %s\n", i+1, row[0])
		ans := row[1]

		input, answered := askQuestion(ques, timeLimit, answerCh)

		if ans == input {
			fmt.Println("Correct!")
			correctCnt++
		} else if answered {
			fmt.Println("Incorrect!")
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correctCnt, N)
}
