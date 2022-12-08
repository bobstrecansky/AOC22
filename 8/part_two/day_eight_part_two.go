package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	forest := readForest(fileScanner)
	res := calculateVisible(forest)
	fmt.Println(res)
}

func readForest(fileScanner *bufio.Scanner) [][]int {
	var forest [][]int
	for fileScanner.Scan() {
		rowString := fileScanner.Text()
		var row []int
		for i := range rowString {
			treeHeight, _ := strconv.Atoi(rowString[i : i+1])
			row = append(row, treeHeight)
		}
		forest = append(forest, row)
	}
	return forest
}

func calculateVisible(forest [][]int) int {
	var res int
	forestLength := len(forest[0])
	forestHeight := len(forest)

	// calculate perimeter (can see all trees on the perimeter)
	res = 2 * (forestHeight + (forestLength - 2))
	isCompletelyVisible := 4
	// calculate the rest of the trees not on the perimeter
	for x := 1; x < forestHeight-1; x++ {
		for y := 1; y < forestLength-1; y++ {
			inspectingTree := forest[x][y]

			// check from the top
			visibleFromTop := 0
			for i := x - 1; i >= 0; i-- {
				visibleFromTop++
				if forest[i][y] >= inspectingTree {
					isCompletelyVisible--
					break
				}
			}

			// check from the bottom
			visibleFromBottom := 0
			for i := x + 1; i < forestHeight; i++ {
				visibleFromBottom++
				if forest[i][y] >= inspectingTree {
					isCompletelyVisible--
					break
				}
			}

			// check from the left
			visibleFromLeft := 0
			for i := y - 1; i >= 0; i-- {
				visibleFromLeft++
				if forest[x][i] >= inspectingTree {
					isCompletelyVisible--
					break
				}
			}

			// check from the right
			visibleFromRight := 0
			for i := y + 1; i < len(forest[0]); i++ {
				visibleFromRight++
				if forest[x][i] >= inspectingTree {
					isCompletelyVisible--
					break
				}
			}

			newHighScore := visibleFromTop * visibleFromBottom * visibleFromLeft * visibleFromRight
			if newHighScore > res {
				res = newHighScore
			}
		}
	}
	return res
}
