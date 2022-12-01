package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed reading File %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var sum, cap int
	for fileScanner.Scan() {
		i, _ := strconv.Atoi(fileScanner.Text())
		sum += i
		if fileScanner.Text() == "" {
			if sum > cap {
				cap = sum
			}
			sum = 0
		}
	}

	fmt.Println("Max Calories: ", cap)
}
