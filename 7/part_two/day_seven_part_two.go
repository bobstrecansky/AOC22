package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	fileScanner.Scan()

	dirSize := map[string]uint64{}
	var currentWorkingDir []string

	// split line into separate string fields
	for fileScanner.Scan() {
		lineStringSplit := strings.Fields(fileScanner.Text())
		if lineStringSplit[0] == "$" {
			if lineStringSplit[1] == "cd" {
				dir := lineStringSplit[2]
				switch dir {
				default:
					currentWorkingDir = append(currentWorkingDir, dir)
				case "/":
					currentWorkingDir = []string{}
				case "..":
					currentWorkingDir = currentWorkingDir[:len(currentWorkingDir)-1]
				}
			}
		} else if lineStringSplit[0] != "dir" {
			currentFileSize, err := strconv.ParseUint(lineStringSplit[0], 10, 64)
			if err != nil {
				log.Fatal("Could not parse current file size")
			}
			for i := len(currentWorkingDir); i >= 0; i-- {
				currentDir := strings.Join(currentWorkingDir[0:i], "/")
				currentDirSize := currentFileSize
				if v, ok := dirSize[currentDir]; ok {
					currentDirSize += v
				}
				dirSize[currentDir] = currentDirSize
			}
		}
		fmt.Println(currentWorkingDir)
	}

	usedDiskSpace := dirSize[""]
	freeDiskSpace := uint64(70000000) - usedDiskSpace
	var res uint64
	for path := range dirSize {
		size := dirSize[path]
		if size+freeDiskSpace >= 30000000 {
			if res == 0 || size < res {
				res = size
			}
		}
	}
	fmt.Println(res)
}
