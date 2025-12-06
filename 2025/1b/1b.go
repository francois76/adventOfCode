package main

import (
	"bufio"
	"strconv"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	shared.Run(func() interface{} {
		val := int64(50)
		count := int64(0)

		shared.Open("../1/1.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			dir := currentLineString[:1]
			currentLine, _ := strconv.ParseInt(currentLineString[1:], 10, 64)

			switch dir {
			case "R":
				val += currentLine
				count += val / 100
				val %= 100
			case "L":
				wasZero := val == 0
				val -= currentLine

				landedOnZero := val%100 == 0
				if landedOnZero {
					count++
				}

				for val < 0 {
					val += 100
					count++
				}

				if wasZero {
					count--
				}
			}

		})

		return count
	})
}
