package main

import (
	"bufio"
	"regexp"
	"strconv"

	"github.com/francois76/adventOfCode/shared"
)

var re = regexp.MustCompile(`(?m)mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

func main() {
	parseInt := func(s string) int64 {
		i, _ := strconv.ParseInt(s, 10, 64)
		return i
	}
	shared.Run(func() interface{} {
		total := int64(0)
		enabled := true
		shared.Open("../3/3.txt", func(fileScanner *bufio.Scanner) {
			for _, match := range re.FindAllStringSubmatch(fileScanner.Text(), -1) {
				if match[0] == "do()" {
					enabled = true
				} else if match[0] == "don't()" {
					enabled = false
				}
				if enabled {
					total += parseInt(match[1]) * parseInt(match[2])
				}

			}
		})

		return total
	})

}
