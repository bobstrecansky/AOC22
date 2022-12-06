package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Failed reading File %s", err)
		}
	}(file)
	if err != nil {
		log.Fatalf("Failed reading File %s", err)
	}
	fileScanner := bufio.NewScanner(file)

	var overlaps int
	for fileScanner.Scan() {
		var firstStart, firstEnd, secondStart, secondEnd int
		fmt.Sscanf(fileScanner.Text(), "%d-%d,%d-%d", &firstStart, &firstEnd, &secondStart, &secondEnd)
		if secondStart <= firstEnd && secondEnd >= firstStart || firstStart <= secondEnd && firstEnd >= secondStart {
			overlaps++
		}
	}
	fmt.Println(overlaps)
}
