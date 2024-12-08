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
		count := int64(0)
		left := []int64{}
		right := []int64{}
		parseInt := func(s string) int64 {
			i, _ := strconv.ParseInt(s, 10, 64)
			return i
		}
		shared.Open("../1/1.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			items := strings.Split(currentLineString, "   ")
			left = append(left, parseInt(items[0]))
			right = append(right, parseInt(items[1]))
		})
		// group items by occurency
		rightGroupedBy := lo.GroupBy(right, func(item int64) int64 {
			return item
		})
		// adds up distance between each list:
		for i := 0; i < len(left); i++ {
			count += left[i] * int64(len(rightGroupedBy[left[i]]))
		}
		return count
	})
}
