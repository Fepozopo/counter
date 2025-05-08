package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "./words.txt"

	log.SetFlags(0)

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read from file: %v", err)
	}

	wordCount := CountWords(data)

	fmt.Println(wordCount)
}

func CountWords(data []byte) int {
	words := bytes.Fields(data)
	return len(words)
}
