package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Read input file
	input, err := os.Open("input.txt")
	defer func(input *os.File) {
		err := input.Close()
		if err != nil {
			log.Fatalf("Failed reading File %s", err)
		}
	}(input)
	if err != nil {
		log.Fatalf("Failed reading File %s", err)
	}
	fileScanner := bufio.NewScanner(input)
	fileScanner.Scan()

	var buffer int = 14

	for i := range fileScanner.Text() {
		chars := make(map[byte]bool)
		for j := 0; j < buffer; j++ {
			chars[fileScanner.Text()[i+j]] = true
		}

		if len(chars) == buffer {
			fmt.Println(i + buffer)
			break
		}
	}
}
