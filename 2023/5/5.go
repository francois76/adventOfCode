package main

import (
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
	"github.com/samber/lo"
)

type rangeStruct struct {
	fromStart int64
	toStart   int64
	size      int64
}

func (r rangeStruct) contains(x int64) bool {
	return x >= r.fromStart && x < r.fromStart+r.size
}

func (r rangeStruct) mapsTo(x int64) int64 {
	return x + r.toStart - r.fromStart
}

func buildRange(input string) []rangeStruct {
	lines := strings.Split(input, "\n")
	return lo.Map(lines, func(line string, _ int) rangeStruct {
		elements := strings.Split(strings.TrimSpace(line), " ")
		toInt := func(s string) int64 {
			i, _ := strconv.ParseInt(s, 10, 64)
			return i
		}
		return rangeStruct{
			toStart:   toInt(elements[0]),
			fromStart: toInt(elements[1]),
			size:      toInt(elements[2]),
		}
	})
}

func main() {
	minPlace := int64(999999999999999999)
	shared.Run(func() any {
		for _, seed := range seeds {
			currentValue := seed
			for i := 0; i <= 6; i++ {
				for _, r := range allRanges[i] {
					if r.contains(currentValue) {
						currentValue = r.mapsTo(currentValue)
						break
					}
				}
			}
			if currentValue < minPlace {
				minPlace = currentValue
			}
		}
		return minPlace
	})
}
