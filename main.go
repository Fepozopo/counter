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

	fmt.Printf("data: %s\n", string(data))
}
