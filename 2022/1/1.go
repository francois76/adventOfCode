package main

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/francois76/adventOfCode/shared"
)

type color string

const yellow color = "yellow"
const blue color = "yellow"
const red color = "red"

func main() {
	var c color
	switch c {
	case yellow:
		fmt.Println("yellow")
	case red:
		fmt.Println("red")
	default:
		fmt.Println("other color")
	}
	shared.Run(func() any {
		maxCount := 0
		currentCount := 0

		shared.Open("1.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			if currentLineString == "" {
				if currentCount > maxCount {
					maxCount = currentCount
				}
				currentCount = 0
			} else {
				currentLine, _ := strconv.ParseInt(currentLineString, 10, 64)
				currentCount += int(currentLine)
			}
		})
		return maxCount
	})
}
