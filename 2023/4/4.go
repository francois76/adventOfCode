package main

import (
	"bufio"
	"strings"

	"github.com/francois76/adventOfCode/shared"
	"github.com/samber/lo"
)

func main() {
	shared.Run(func() interface{} {
		count := int64(0)

		shared.Open("4.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			lineParts := strings.Split(currentLineString, ":")
			// cardNumber := strings.Split(lineParts[0], " ")[1]
			numbers := strings.Split(lineParts[1], "|")
			winningNumbers := strings.Split(numbers[0], " ")
			obtainedNumber := strings.Split(numbers[1], " ")
			lineCount := lo.Reduce(obtainedNumber, func(acc int64, number string, _ int) int64 {
				if number == "" {
					return acc
				}
				if lo.Contains(winningNumbers, number) {
					if acc == 0 {
						return 1
					}
					return acc * 2
				}
				return acc
			}, 0)
			count += lineCount
		})
		return count
	})
}
