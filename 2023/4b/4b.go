package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
	"github.com/samber/lo"
)

func main() {
	gottenCard := map[int64]int64{}
	for i := int64(1); i <= 213; i++ {
		gottenCard[i] = 0
	}
	shared.Run(func() any {

		shared.Open("../4/4.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			lineParts := strings.Split(currentLineString, ":")
			cardNumber, _ := strconv.ParseInt(lo.Filter(strings.Split(lineParts[0], " "), func(item string, index int) bool {
				return item != ""
			})[1], 10, 64)
			fmt.Println("cardNumber: ", cardNumber)
			gottenCard[cardNumber]++ // adding the original card
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
					return acc + 1
				}
				return acc
			}, 0)
			for k := int64(1); k <= gottenCard[cardNumber]; k++ {
				for i := int64(cardNumber + 1); i <= int64(cardNumber+lineCount); i++ {
					gottenCard[i]++
				}
			}

		})
		count := int64(0)
		for i := int64(1); i <= 213; i++ {
			count += gottenCard[i]
		}
		return count
	})
}
