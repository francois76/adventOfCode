package main

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	twoDigitsMatcher, _ := regexp.Compile(`(?m)\D*(\d).*(\d)\D*`)
	oneDigitsMatcher, _ := regexp.Compile(`(?m)\D*(\d)\D*`)
	shared.Run(func() interface{} {
		count := int64(0)

		shared.Open("../1/1.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := transformLine(fileScanner.Text())
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

func transformLine(input string) string {
	output := strings.ReplaceAll(input, "sevenine", "79")
	output = strings.ReplaceAll(output, "oneight", "18")
	output = strings.ReplaceAll(output, "threeight", "38")
	output = strings.ReplaceAll(output, "fiveight", "58")
	output = strings.ReplaceAll(output, "nineight", "98")
	output = strings.ReplaceAll(output, "eighthree", "83")
	output = strings.ReplaceAll(output, "eightwo", "82")
	output = strings.ReplaceAll(output, "twone", "21")
	output = strings.ReplaceAll(output, "one", "1")
	output = strings.ReplaceAll(output, "two", "2")
	output = strings.ReplaceAll(output, "three", "3")
	output = strings.ReplaceAll(output, "four", "4")
	output = strings.ReplaceAll(output, "five", "5")
	output = strings.ReplaceAll(output, "six", "6")
	output = strings.ReplaceAll(output, "seven", "7")
	output = strings.ReplaceAll(output, "eight", "8")
	output = strings.ReplaceAll(output, "nine", "9")
	return output
}
