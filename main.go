package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("./words.txt")
	if err != nil {
		log.Fatalf("failed to read from file: %v", err)
	}

	wordCount := 0
	for _, v := range data {
		if v == ' ' {
			wordCount++
		}
	}
	wordCount++

	fmt.Println(wordCount)
}
