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

func (s *stack) push(r rune) {
	s.items = append(s.items, r)
}

func (s *stack) pop() (r rune) {
	r = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return
}

func (s *stack) appendToBottom(r rune) {
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
				stacks[k/4].appendToBottom(v)
			}
		}
		fileScanner.Scan()
	}

	// Empty line between table and directives
	fileScanner.Scan()

	for fileScanner.Scan() {
		var toMove, from, to int
		_, err := fmt.Sscanf(fileScanner.Text(), "move %d from %d to %d", &toMove, &from, &to)
		if err != nil {
			return
		}
		for move := 0; move < toMove; move++ {
			stacks[to-1].push(stacks[from-1].pop())
		}
	}
	for _, s := range stacks {
		fmt.Print(string(s.pop()))
	}
}
