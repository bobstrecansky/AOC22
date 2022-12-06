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

	var contained int
	for fileScanner.Scan() {
		var firstBegin, firstEnd, secondStart, secondEnd int
		fmt.Sscanf(fileScanner.Text(), "%d-%d,%d-%d", &firstBegin, &firstEnd, &secondStart, &secondEnd)
		if secondStart >= firstBegin && secondEnd <= firstEnd || firstBegin >= secondStart && firstEnd <= secondEnd {
			contained++
		}
	}
	fmt.Println(contained)
}
