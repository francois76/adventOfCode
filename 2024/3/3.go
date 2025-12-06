package main

import (
	"bufio"
	"regexp"
	"strconv"

	"github.com/francois76/adventOfCode/shared"
)

var re = regexp.MustCompile(`(?m)mul\((\d+),(\d+)\)`)

func main() {
	parseInt := func(s string) int64 {
		i, _ := strconv.ParseInt(s, 10, 64)
		return i
	}
	shared.Run(func() any {
		total := int64(0)
		shared.Open("3.txt", func(fileScanner *bufio.Scanner) {
			for _, match := range re.FindAllStringSubmatch(fileScanner.Text(), -1) {
				total += parseInt(match[1]) * parseInt(match[2])
			}
		})

		return total
	})

}
