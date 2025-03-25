package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

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

	for i, row := range data {
		ques := row[0]
		fmt.Printf("Problem #%d: %s\n", i+1, ques)

		ans, err := strconv.Atoi(row[1])
		if err != nil {
			panic(err)
		}

		var input int
		fmt.Scanf("%d", &input)

		if ans == input {
			fmt.Println("Correct!")
			correctCnt++
		} else {
			fmt.Println("Incorrect!")
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correctCnt, N)
}
