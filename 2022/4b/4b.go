package main

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	shared.Run(func() interface{} {
		count := 0
		shared.Open("../4/4.txt", func(fileScanner *bufio.Scanner) {
			count += hasOneFullyContains(fileScanner.Text())
		})
		return count
	})
}

func hasOneFullyContains(line string) int {
	ranges := strings.Split(line, ",")
	leftRange := strings.Split(ranges[0], "-")
	rightRange := strings.Split(ranges[1], "-")
	leftRangeMin, _ := strconv.ParseInt(leftRange[0], 10, 64)
	leftRangeMax, _ := strconv.ParseInt(leftRange[1], 10, 64)
	rightRangeMin, _ := strconv.ParseInt(rightRange[0], 10, 64)
	rightRangeMax, _ := strconv.ParseInt(rightRange[1], 10, 64)
	if rightRangeMax < leftRangeMin || leftRangeMax < rightRangeMin {
		return 0
	}
	return 1
}
