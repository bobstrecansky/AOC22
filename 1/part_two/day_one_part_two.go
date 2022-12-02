package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed reading File %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var sum, res int
	var top []int
	for fileScanner.Scan() {
		i, _ := strconv.Atoi(fileScanner.Text())
		sum += i
		if fileScanner.Text() == "" {
			top = append(top, sum)
			sum = 0
		}
	}

	sort.Ints(top)
	for i := 1; i < 4; i++ {
		res += top[len(top)-i]
	}
	fmt.Println("Sum of Top 3 Calories", res)
}
