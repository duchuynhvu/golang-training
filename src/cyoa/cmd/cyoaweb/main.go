package main

import (
	"fmt"
	"flag"
	"os"
	"cyoa"
)

func main() {
	file := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()

	fmt.Printf("Using the story in %s.\n", *file)
	
	f, err := os.Open(*file)
	if err != nil {
		fmt.Println("Failed to open the JSON provided file")
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}