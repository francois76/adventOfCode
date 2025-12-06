package main

import (
	"fmt"
	"math"
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

func (r rangeStruct) subRange(x basicRange) (basicRange, bool) {
	if x.end() <= r.fromStart || x.begin >= r.fromStart+r.size {
		return basicRange{}, false
	}
	result := basicRange{
		begin:  int64(math.Max(float64(x.begin), float64(r.fromStart))),
		length: int64(math.Min(float64(x.end()), float64(r.fromStart+r.size))) - int64(math.Max(float64(x.begin), float64(r.fromStart))),
	}
	return result, true
}

func (r rangeStruct) mapsTo(x basicRange) basicRange {
	return basicRange{
		begin:  x.begin + r.toStart - r.fromStart,
		length: x.end() + r.toStart - r.fromStart - (x.begin + r.toStart - r.fromStart),
	}
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
			currentValue := []basicRange{seed}
			for i := 0; i <= 6; i++ {
				currentValueStep := []basicRange{}
				for _, r := range allRanges[i] {
					for _, c := range currentValue {
						if subRange, ok := r.subRange(c); ok {
							mapsStep := r.mapsTo(subRange)
							currentValueStep = append(currentValueStep, mapsStep)
						} else {
							currentValueStep = append(currentValueStep, c)
						}
					}
				}
				fmt.Println(currentValueStep)
				copy(currentValue, currentValueStep)
			}

		}
		return minPlace
	})
}
