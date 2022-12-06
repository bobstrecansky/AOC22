package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
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

	var prioritySum int

	for fileScanner.Scan() {
		items := make(map[rune]bool)
		for _, rucksackLeft := range fileScanner.Text()[:len(fileScanner.Text())/2] {
			items[rucksackLeft] = true
		}

		for _, rucksackRight := range fileScanner.Text()[len(fileScanner.Text())/2:] {
			if items[rucksackRight] {
				prioritySum += int(unicode.ToLower(rucksackRight) - 96)
				if unicode.IsUpper(rucksackRight) {
					prioritySum += 26
				}
				break
			}
		}
	}
	fmt.Println(prioritySum)
}
