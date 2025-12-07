package main

import (
	"bufio"
	"fmt"
	"regexp"

	"github.com/francois76/adventOfCode/shared"
)

func buildRegexStepOne(digit int) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf("(?m)[^%d]*%d(.+)", digit, digit))
}

func buildRegexStepTwo(digit int) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf("(?m)[^%d]*%d(.*)", digit, digit))
}

var regexesStep1 = map[int]*regexp.Regexp{
	9: buildRegexStepOne(9),
	8: buildRegexStepOne(8),
	7: buildRegexStepOne(7),
	6: buildRegexStepOne(6),
	5: buildRegexStepOne(5),
	4: buildRegexStepOne(4),
	3: buildRegexStepOne(3),
	2: buildRegexStepOne(2),
	1: buildRegexStepOne(1),
	0: buildRegexStepOne(0),
}

var regexesStep2 = map[int]*regexp.Regexp{
	9: buildRegexStepTwo(9),
	8: buildRegexStepTwo(8),
	7: buildRegexStepTwo(7),
	6: buildRegexStepTwo(6),
	5: buildRegexStepTwo(5),
	4: buildRegexStepTwo(4),
	3: buildRegexStepTwo(3),
	2: buildRegexStepTwo(2),
	1: buildRegexStepTwo(1),
	0: buildRegexStepTwo(0),
}

func main() {
	shared.Run(func() any {
		count := int64(0)
		shared.Open("3.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			power := 0
			for i := 9; i >= 0; i-- {
				subst := regexesStep1[i].ReplaceAllString(currentLineString, "$1")
				if subst != currentLineString {
					power = i * 10
					currentLineString = subst
					break
				}
			}
			fmt.Println(currentLineString)
			for i := 9; i >= 0; i-- {
				subst := regexesStep2[i].ReplaceAllString(currentLineString, "$1")
				if subst != currentLineString {
					power = power + i
					break
				}
			}
			count = count + int64(power)

		})

		return count
	})
}
