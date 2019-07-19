package main

import (
	"fmt"
	"flag"
	"os"
	"encoding/csv"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse the provided CSV file."))
	}
	

}

type problem struct {
	q string
	a string
}


func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}