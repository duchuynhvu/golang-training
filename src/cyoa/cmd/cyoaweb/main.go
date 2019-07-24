package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
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

	h := cyoa.NewHandler(story, nil)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
