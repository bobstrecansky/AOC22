package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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

		fmt.Println("SCORE:", score, " OPP:", opp, " USER:", user)

		// Rock, +1 point
		if user == "X" {
			score += 1
		}

		// Paper, +2 points
		if user == "Y" {
			score += 2
		}

		// Scissors, +3 points
		if user == "Z" {
			score += 3
		}

		// Tie conditions, +3 points
		if opp == "A" && user == "X" || opp == "B" && user == "Y" || opp == "C" && user == "Z" {
			score += 3
		}

		// Winning conditions, +6 points
		// Elf throws rock, User throws paper
		// Elf throws paper, User throws scissors
		// Elf throws scissors, User throws rock
		if opp == "A" && user == "Y" || opp == "B" && user == "Z" || opp == "C" && user == "X" {
			score += 6
		}

		// if we lose, we are granted 0 points (no conditional necessary)
		fmt.Println(score)
	}
}
