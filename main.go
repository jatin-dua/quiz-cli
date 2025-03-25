package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
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

	for _, row := range data {
		ques, ans := row[0], row[1]
		fmt.Printf("%s, %s\n", ques, ans)
	}
}
