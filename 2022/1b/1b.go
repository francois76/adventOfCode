package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("../1/1.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	maxCount := make([]int, 3)
	currentCount := 0

	for fileScanner.Scan() {
		currentLineString := fileScanner.Text()
		if currentLineString == "" {
			if currentCount >= maxCount[2] {
				maxCount = []int{maxCount[1], maxCount[2], currentCount}
			}
			currentCount = 0
		} else {
			currentLine, _ := strconv.ParseInt(currentLineString, 10, 64)
			currentCount += int(currentLine)
		}
	}
	fmt.Println(maxCount[0] + maxCount[1] + maxCount[2])
	readFile.Close()
}
