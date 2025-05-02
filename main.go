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

	wordCount := CountWords(data)

	fmt.Println(wordCount)
}

func CountWords(data []byte) int {
	if len(data) == 0 {
		return 0
	}

	wordCount := 0
	for _, v := range data {
		if v == ' ' {
			wordCount++
		}
	}
	wordCount++

	return wordCount
}
