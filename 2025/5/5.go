package main

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	shared.Run(func() any {
		count := 0
		hasComputedRanges := false
		allRanges := []ranges{}
		shared.Open("5.txt", func(fileScanner *bufio.Scanner) {
			line := fileScanner.Text()
			if line == "" {
				hasComputedRanges = true
			}
			if !hasComputedRanges {
				allRanges = append(allRanges, buildRange(line))
			} else {
				for _, r := range allRanges {
					value, _ := strconv.ParseInt(line, 10, 64)
					if r.contains(value) {
						count++
						break
					}
				}
			}
		})
		return count
	})
}

type ranges struct {
	min int64
	max int64
}

func (r *ranges) contains(value int64) bool {
	return value >= r.min && value <= r.max
}

func buildRange(line string) ranges {
	items := strings.Split(line, "-")
	min, _ := strconv.ParseInt(items[0], 10, 64)
	max, _ := strconv.ParseInt(items[1], 10, 64)

	return ranges{min: min, max: max}
}
