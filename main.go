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

	wordCount := countWords(data)

	fmt.Println(wordCount)
}

func countWords(data []byte) int {
	wordCount := 0
	for _, v := range data {
		if v == ' ' {
			wordCount++
		}
	}
	wordCount++

	return wordCount
}
