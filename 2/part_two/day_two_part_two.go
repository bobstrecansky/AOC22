package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// main - day two part two
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed reading File %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var score int
	for fileScanner.Scan() {
		f := strings.Split(fileScanner.Text(), " ")
		opp := f[0]
		user := f[1]

		// Need to lose
		if user == "X" {
			// opp is rock, we need scissors
			if opp == "A" {
				score += 3
			}

			// opp is paper, we need rock
			if opp == "B" {
				score++
			}

			// opp is scissors, we need paper
			if opp == "C" {
				score += 2
			}
		}

		// Need to tie
		if user == "Y" {
			score += 3

			if opp == "A" {
				score++
			}
			if opp == "B" {
				score += 2
			}
			if opp == "C" {
				score += 3
			}
		}

		// Need to win
		if user == "Z" {
			score += 6

			// opp is rock, we need paper
			if opp == "A" {
				score += 2
			}

			// opp is paper, we need scissors
			if opp == "B" {
				score += 3
			}

			// opp is scissors, we need rock
			if opp == "C" {
				score++
			}
		}
	}

	fmt.Println(score)
}
