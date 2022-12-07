package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

func main() {

	shared.Open("6.txt", func(fileScanner *bufio.Scanner) {
		characters := strings.Split(fileScanner.Text(), "")
		for idx := 0; idx < len(characters)-4; idx++ {
			currentFrame := characters[idx : idx+4]
			set := map[string]bool{}
			for _, letter := range currentFrame {
				set[letter] = true
			}
			if len(set) == 4 {
				fmt.Println(idx + 4)
				return
			}
		}
	})
}
