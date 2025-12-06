package main

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
	"github.com/samber/lo"
)

func main() {
	shared.Run(func() any {

		total := int64(0)
		lineRegex, _ := regexp.Compile(`(?m)Game (\d+):(.+)`)
		colorMatcher, _ := regexp.Compile(`( (\d+) (red))|( (\d+) (green))|( (\d+) (blue))`)

		shared.Open("../2/2.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			elements := lineRegex.FindAllStringSubmatch(currentLineString, -1)
			lines := strings.Split(elements[0][2], ";")
			valuesList := []map[string]string{}
			for line := range lines {
				colors := colorMatcher.FindAllStringSubmatch(lines[line], -1)
				values := map[string]string{}
				for i := range colors {
					values[colors[i][3]] = colors[i][2]
					values[colors[i][6]] = colors[i][5]
					values[colors[i][9]] = colors[i][8]
				}
				valuesList = append(valuesList, values)
			}
			total += computePower(valuesList)
		})
		return total
	})
}

func computePower(valuesList []map[string]string) int64 {
	accumulator := func(color string) int64 {
		return lo.Reduce(valuesList, func(acc int64, values map[string]string, _ int) int64 {
			colorAmount, _ := strconv.ParseFloat(values[color], 64)
			if colorAmount == 0 {
				return acc
			}
			return lo.Max([]int64{acc, int64(colorAmount)})
		}, 0)
	}
	return accumulator("red") * accumulator("green") * accumulator("blue")
}
