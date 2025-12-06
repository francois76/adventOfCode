package main

import (
	"bufio"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	shared.Run(func() any {

		shared.Open("NUMBER.txt", func(fileScanner *bufio.Scanner) {
			fileScanner.Text()
		})
		return 0
	})
}
