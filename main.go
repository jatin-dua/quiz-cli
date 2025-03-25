package main

import (
	"flag"
	"fmt"
)

func main() {
	var filename string
	var limit int

	flag.StringVar(&filename, "csv", "problems.csv", "a csv file in the format of 'question, answer' (default 'problems.csv')")
	flag.IntVar(&limit, "limit", 30, "the time limit for quiz in seconds (default 30)")
	flag.Parse()

	fmt.Printf("Filename arg: %s\n", filename)
	fmt.Printf("Limit arg: %d\n", limit)
}
