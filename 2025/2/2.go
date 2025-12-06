package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	buildPairNumber := func(n int64) int64 {
		result, _ := strconv.ParseInt(fmt.Sprintf("%d%d", n, n), 10, 64)
		return result
	}
	shared.Run(func() any {
		count := int64(0)

		shared.Open("2.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			items := strings.Split(currentLineString, "-")
			start, _ := strconv.ParseInt(items[0], 10, 64)
			startStart, _ := strconv.ParseInt(items[0][0:len(items[0])/2], 10, 64)
			end, _ := strconv.ParseInt(items[1], 10, 64)
			for {
				i := buildPairNumber(startStart)
				fmt.Println(i)
				if i < start {
					// could happen if begin part smaller than end part
					startStart++
					continue
				}
				if i > end {
					break
				}
				count += i
				startStart++
			}

		})

		return count
	})
}
