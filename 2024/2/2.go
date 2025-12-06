package main

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
	"github.com/samber/lo"
)

func main() {
	shared.Run(func() any {
		parseInt := func(s string) int64 {
			i, _ := strconv.ParseInt(s, 10, 64)
			return i
		}
		countSafe := 0
		shared.Open("2.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			vector := strings.Split(currentLineString, " ")
			intVector := lo.Map(vector, func(item string, _ int) int64 {
				return parseInt(item)
			})
			safe, increasing := isSafeZero(intVector)
			if !safe {
				return
			}
			countSafe += func() int {
				if isSafe(intVector, 1, increasing) {
					return 1
				}
				return 0
			}()
		})
		return countSafe
	})
}

func isSafeZero(vector []int64) (safe bool, increasing bool) {
	diff := vector[1] - vector[0]
	if checkDelta(diff) {
		return true, true
	} else if checkDelta(-diff) {
		return true, false
	}
	return false, false

}

func isSafe(vector []int64, n int, increasing bool) bool {
	// if we reached the end of the array, we are safe
	if len(vector)-1 <= n {
		return true
	}
	diff := vector[n+1] - vector[n]
	if !increasing {
		diff = diff * -1
	}
	if checkDelta(diff) {
		return isSafe(vector, n+1, increasing)
	}
	return false
}

func checkDelta(diff int64) bool {
	return diff == 1 || diff == 2 || diff == 3
}
