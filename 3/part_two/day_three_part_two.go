package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
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

	var sumOfPriorities int

	for fileScanner.Scan() {
		firstRuckItem := itemSet(fileScanner.Text())
		fileScanner.Scan()
		secondRuckItem := itemSet(fileScanner.Text())
		fileScanner.Scan()
		thirdRuckItem := itemSet(fileScanner.Text())

		for firstElfItem := range firstRuckItem {
			if secondRuckItem[firstElfItem] && thirdRuckItem[firstElfItem] {
				sumOfPriorities += int(unicode.ToLower(firstElfItem) - 96)
				if unicode.IsUpper(firstElfItem) {
					sumOfPriorities += 26
				}
				break
			}
		}
	}
	fmt.Println(sumOfPriorities)
}

func itemSet(items string) (res map[rune]bool) {
	res = make(map[rune]bool)
	for _, item := range items {
		res[item] = true
	}
	return res
}
