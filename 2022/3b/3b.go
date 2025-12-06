package main

import (
	"bufio"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	shared.Run(func() any {

		// variables
		m := make(map[string]int)

		// Populate the map with keys and values
		for i, ch := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
			m[string(ch)+string(ch)+string(ch)] = i + 1
		}

		totalCount := 0
		shared.Open("../3/3.txt", func(fileScanner *bufio.Scanner) {
			firstPart := fileScanner.Text()
			fileScanner.Scan()
			secondPart := fileScanner.Text()
			fileScanner.Scan()
			totalCount += getCommonChar(m, firstPart, secondPart, fileScanner.Text())
		})
		return totalCount
	})
}

func getCommonChar(m map[string]int, firstPart string, secondPart string, thirdPart string) int {
	for _, x := range firstPart {
		for _, y := range secondPart {
			for _, z := range thirdPart {
				if rs := m[string(x)+string(y)+string(z)]; rs != 0 {
					return rs
				}
			}
		}
	}
	return 0
}
