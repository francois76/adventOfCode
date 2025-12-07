package main

import (
	"bufio"
	"fmt"
	"math"
	"regexp"

	"github.com/francois76/adventOfCode/shared"
)

func buildRegexStep(digit int, step int) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf("(?m)[^%d]*%d(.{%d,})", digit, digit, step-1))
}

func main() {
	shared.Run(func() any {
		count := int64(0)
		shared.Open("../3/3.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			power := 0
			for step := 12; step >= 0; step-- {
				for i := 9; i >= 0; i-- {
					subst := buildRegexStep(i, step).ReplaceAllString(currentLineString, "$1")
					if subst != currentLineString {
						power += i * int(math.Pow10(step-1))
						currentLineString = subst
						fmt.Println("for step ", step, "taken digit ", i, " power is now ", power, "remaining parti is now ", currentLineString)
						break
					}
				}
			}
			fmt.Println(power)

			count = count + int64(power)

		})

		return count
	})
}
