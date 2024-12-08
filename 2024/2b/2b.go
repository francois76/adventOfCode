package main

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
	"github.com/samber/lo"
)

func main() {
	shared.Run(func() interface{} {
		parseInt := func(s string) int64 {
			i, _ := strconv.ParseInt(s, 10, 64)
			return i
		}
		countSafe := 0
		shared.Open("../2/2.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			vector := strings.Split(currentLineString, " ")
			intVector := lo.Map(vector, func(item string, _ int) int64 {
				return parseInt(item)
			})
			oneOrientation := func(assumeIncreasing bool) bool {
				return isSafe(intVector, 0, assumeIncreasing, true) || isSafe(intVector, 1, assumeIncreasing, false)
			}

			countSafe += func() int {
				if oneOrientation(true) || oneOrientation(false) {
					return 1
				}
				return 0
			}()
		})
		return countSafe
	})
}

func isSafe(vector []int64, n int, assumeIncreasing bool, hasJoker bool) bool {
	// if we reached the end of the array, we are safe
	if len(vector)-1 <= n {
		return true
	}
	diff := vector[n+1] - vector[n]
	if !assumeIncreasing {
		diff = diff * -1
	}
	if checkDelta(diff) {
		return isSafe(vector, n+1, assumeIncreasing, hasJoker)
	}
	if !hasJoker {
		// checkdelta is wrong with given orientation so its wrong
		return false
	}
	// if we are one step of the end and had our joker, we are right
	if len(vector)-1 <= n+1 {
		return true
	}
	diff2 := vector[n+2] - vector[n]
	if !assumeIncreasing {
		diff2 = diff2 * -1
	}
	if checkDelta(diff2) {
		return isSafe(vector, n+2, assumeIncreasing, false)
	}
	return false
}

func checkDelta(diff int64) bool {
	return diff == 1 || diff == 2 || diff == 3
}
