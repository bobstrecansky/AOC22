package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Read input file
	input, err := os.Open("sample.txt")
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

	// create a map of x-y coordinates (stored as string) to bool
	tailVisited := map[string]bool{"0 0": true}
	currentLocation := "0 0"

	for fileScanner.Scan() {
		movement := strings.Fields(fileScanner.Text())
		direction := movement[0]
		distance, _ := strconv.Atoi(movement[1])
		currentLocation = calculateDirection(tailVisited, currentLocation, direction, distance)
	}

	PrettyPrint(tailVisited)
}

func calculateDirection(tailVisited map[string]bool, currentLocation, direction string, distance int) string {
	switch direction {
	case "U":
		tmp := strings.Fields(currentLocation)
		y, _ := strconv.Atoi(tmp[1])
		for i := 0; i < y+distance; i++ {
			tailVisited[fmt.Sprintf("%v %v", tmp[0], y+i)] = true
		}
		res := fmt.Sprintf("%v %v", tmp[0], y+distance)
		fmt.Println("U:", res)
		return res
	case "D":
		tmp := strings.Fields(currentLocation)
		y, _ := strconv.Atoi(tmp[1])
		for i := y - distance; i > distance; i-- {
			tailVisited[fmt.Sprintf("%v %v", tmp[0], y-i)] = true
		}
		res := fmt.Sprintf("%v %v", tmp[0], y-distance)
		fmt.Println("D:", res)
		return res
	case "R":
		tmp := strings.Fields(currentLocation)
		x, _ := strconv.Atoi(tmp[0])
		for i := 0; i < x+distance; i++ {
			tailVisited[fmt.Sprintf("%v %v", x+i, tmp[1])] = true
		}
		res := fmt.Sprintf("%v %v", x+distance, tmp[1])
		fmt.Println("R:", res)
		return res
	case "L":
		tmp := strings.Fields(currentLocation)
		x, _ := strconv.Atoi(tmp[0])
		for i := x - distance; i > distance; i-- {
			fmt.Println("x-i: ", x-i)
			tailVisited[fmt.Sprintf("%v %v", x-i, tmp[1])] = true
		}
		res := fmt.Sprintf("%v %v", x-distance, tmp[1])
		fmt.Println("L:", res)
		return res
	}
	return "unknown coordinates"
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
