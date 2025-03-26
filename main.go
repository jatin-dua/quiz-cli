package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Problem struct {
	question string
	answer   int
}

func readUserInput(answerCh chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		answerCh <- scanner.Text()
	}
	close(answerCh)
}

func generateProblem(problemId, maxn int) Problem {
	rand.Seed(time.Now().UnixNano())
	num1 := rand.Intn(maxn) + 1
	num2 := rand.Intn(maxn) + 1

	operators := []string{"+", "-", "*", "/"}
	op := operators[rand.Intn(len(operators))]

	var result int
	question := ""

	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}

	question = fmt.Sprintf("Problem #%d: %d %s %d = ?", problemId, num1, op, num2)
	return Problem{question: question, answer: result}
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
	var limit int
	var maxn int

	flag.IntVar(&maxn, "maxn", 100, "the maximum number in a problem (default 100)")
	flag.IntVar(&limit, "limit", 30, "the time limit for quiz in seconds (default 30)")
	flag.Parse()

	correctCnt := 0

	timeLimit := time.Duration(limit) * time.Second
	answerCh := make(chan string, 1)

	go readUserInput(answerCh)

	i := 1
	for {
		problem := generateProblem(i, maxn)
		input, answered := askQuestion(problem.question, timeLimit, answerCh)

		if input == "" && answered {
			break
		}

		if !answered {
			continue
		}

		userInput, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}

		if problem.answer == userInput {
			fmt.Println("Correct!")
			correctCnt++
		} else {
			fmt.Println("Incorrect!")
		}

		i++
	}

	fmt.Printf("You scored %d out of %d.\n", correctCnt, i)
}
