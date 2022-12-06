package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type stack struct {
	items []rune
}

func (s *stack) push(r []rune) {
	s.items = append(s.items, r...)
}

func (s *stack) pop(n int) (r []rune) {
	r = s.items[len(s.items)-n : len(s.items)]
	s.items = s.items[:len(s.items)-n]
	return
}

func (s *stack) addToBottom(r rune) {
	s.items = append([]rune{r}, s.items...)
}

func (s stack) String() string {
	var str string
	for _, r := range s.items {
		str += string(r) + " "
	}
	return str
}

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

	// making for max width
	stacks := make([]stack, 9)
	for !strings.Contains(fileScanner.Text(), "1") {
		for k, v := range fileScanner.Text() {
			if v != ' ' && v != '[' && v != ']' {
				stacks[k/4].addToBottom(v)
			}
		}
		fileScanner.Scan()
	}

	// Empty line between table and directives
	fileScanner.Scan()

	for fileScanner.Scan() {
		var toMove, from, to int
		fmt.Sscanf(fileScanner.Text(), "move %d from %d to %d", &toMove, &from, &to)

		stacks[to-1].push(stacks[from-1].pop(toMove))
	}
	for _, s := range stacks {
		fmt.Print(string(s.pop(1)))
	}
}
