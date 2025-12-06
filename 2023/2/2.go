package main

import (
	"bufio"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	shared.Run(func() any {

		total := 0
		lineRegex, _ := regexp.Compile(`(?m)Game (\d+):(.+)`)
		colorMatcher, _ := regexp.Compile(`( (\d+) (red))|( (\d+) (green))|( (\d+) (blue))`)

		shared.Open("2.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			elements := lineRegex.FindAllStringSubmatch(currentLineString, -1)
			id, _ := strconv.ParseInt(elements[0][1], 10, 64)
			lines := strings.Split(elements[0][2], ";")
			valid := true
			for line := range lines {
				colors := colorMatcher.FindAllStringSubmatch(lines[line], -1)
				values := map[string]string{}
				for i := range colors {
					values[colors[i][3]] = colors[i][2]
					values[colors[i][6]] = colors[i][5]
					values[colors[i][9]] = colors[i][8]
				}
				valid = valid && isValid(values)
			}
			if valid {
				total += int(id)
			}

		})
		return total
	})
}

func isValid(input map[string]string) bool {
	red, _ := strconv.ParseFloat(input["red"], 64)
	green, _ := strconv.ParseFloat(input["green"], 64)
	blue, _ := strconv.ParseFloat(input["blue"], 64)
	return math.Max(0, 13-red)*math.Max(0, 14-green)*math.Max(0, 15-blue) != 0
}
