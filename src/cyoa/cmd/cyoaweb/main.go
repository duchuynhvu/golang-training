package main

import (
	"fmt"
	"flag"
	"os"
	"cyoa"
	"encoding/json"
)

func main() {
	file := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()

	fmt.Printf("Using the story in %s.\n", *file)
	
	f, err := os.Open(*file)
	if err != nil {
		fmt.Println("Failed to open the JSON provided file")
	}

	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		fmt.Println("Failed to decode the JSON file")
	}

	fmt.Printf("%+v\n", story)
}