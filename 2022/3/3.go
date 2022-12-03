package main

import (
	"bufio"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	shared.Run(func() interface{} {

		// variables
		m := make(map[string]int)

		// Populate the map with keys and values
		for i, ch := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
			m[string(ch)+string(ch)] = i + 1
		}

		totalCount := 0

		shared.Open("3.txt", func(fileScanner *bufio.Scanner) {
			totalCount += getCommonChar(m, fileScanner.Text())
		})

		return totalCount
	})
}

func getCommonChar(m map[string]int, line string) int {
	firstPart := line[:len(line)/2]
	secondpart := line[len(line)/2:]
	for _, x := range firstPart {
		for _, y := range secondpart {
			if rs := m[string(x)+string(y)]; rs != 0 {
				return rs
			}
		}
	}
	return 0
}
