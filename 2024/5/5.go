package main

import (
	"bufio"
	"slices"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	rules := []string{}
	parseInt := func(s string) int64 {
		i, _ := strconv.ParseInt(s, 10, 64)
		return i
	}
	isSorted := func(input []string, rules []string) bool {
		original := slices.Clone(input)
		slices.SortFunc(input, func(a, b string) int {
			if a == b {
				return 0
			}
			// items are considered sorted by default except if there is a|b in rules
			data := b + "|" + a
			if slices.Contains(rules, data) {
				return 1
			}
			return -1
		})
		return slices.Equal(input, original)
	}
	itemOfTheMiddle := func(input []string) int64 {
		return parseInt(input[len(input)/2])
	}
	counter := int64(0)
	shared.Run(func() any {
		shared.Open("5.txt", func(fileScanner *bufio.Scanner) {
			line := fileScanner.Text()
			if strings.Contains(line, "|") {
				rules = append(rules, line)
			} else if line != "" {
				lineToCheck := strings.Split(line, ",")
				if isSorted(lineToCheck, rules) {
					counter += itemOfTheMiddle(lineToCheck)
				}

			}
		})
		return counter
	})
}
