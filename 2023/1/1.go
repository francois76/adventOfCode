package main

import (
	"bufio"
	"regexp"
	"strconv"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	twoDigitsMatcher, _ := regexp.Compile(`(?m)\D*(\d).*(\d)\D*`)
	oneDigitsMatcher, _ := regexp.Compile(`(?m)\D*(\d)\D*`)
	shared.Run(func() interface{} {
		count := int64(0)

		shared.Open("1.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			if matches := twoDigitsMatcher.FindAllStringSubmatch(currentLineString, -1); len(matches) == 1 {
				number1, _ := strconv.ParseInt(matches[0][1]+matches[0][2], 10, 64)
				count += number1
			} else if matches := oneDigitsMatcher.FindAllStringSubmatch(currentLineString, -1); len(matches) == 1 {
				number1, _ := strconv.ParseInt(matches[0][1]+matches[0][1], 10, 64)
				count += number1
			} else {
				panic("error with " + currentLineString)
			}
		})
		return count
	})
}
