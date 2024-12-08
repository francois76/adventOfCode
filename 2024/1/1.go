package main

import (
	"bufio"
	"slices"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	shared.Run(func() interface{} {
		count := int64(0)
		left := []int64{}
		right := []int64{}
		parseInt := func(s string) int64 {
			i, _ := strconv.ParseInt(s, 10, 64)
			return i
		}
		shared.Open("1.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			items := strings.Split(currentLineString, "   ")
			left = append(left, parseInt(items[0]))
			right = append(right, parseInt(items[1]))
		})
		slices.Sort(left)
		slices.Sort(right)
		// adds up distance between each list:
		for i := 0; i < len(left); i++ {
			toAdd := (left[i] - right[i])
			if toAdd < 0 {
				toAdd *= -1
			}
			count += toAdd
		}
		return count
	})
}
