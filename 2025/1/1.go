package main

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	shared.Run(func() interface{} {
		dial := int64(50)
		count := 0

		shared.Open("1.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			currentLineString = strings.ReplaceAll(currentLineString, "L", "-")
			currentLineString = strings.ReplaceAll(currentLineString, "R", "")
			currentLine, _ := strconv.ParseInt(currentLineString, 10, 64)
			dial += currentLine
			dial = dial % 100
			if dial == 0 {
				count++
			}
		})

		return count
	})
}
