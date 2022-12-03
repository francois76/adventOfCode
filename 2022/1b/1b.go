package main

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	maxCount := make([]int, 3)
	currentCount := 0

	shared.Open("../1/1.txt", func(fileScanner *bufio.Scanner) {
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
	})

	fmt.Println(maxCount[0] + maxCount[1] + maxCount[2])
}
