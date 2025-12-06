package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
	"github.com/samber/lo"
)

func main() {
	shared.Run(func() any {
		var stones []int64
		shared.Open("11.txt", func(fileScanner *bufio.Scanner) {
			stones = lo.Map(strings.Split(fileScanner.Text(), " "), func(item string, _ int) int64 {
				return parseInt(item)
			})
		})
		for range 25 {
			stones = blink(stones)
		}
		return len(stones)
	})
}

func parseInt(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func blink(input []int64) []int64 {
	output := []int64{}
	for _, stone := range input {
		lengthStone := len(fmt.Sprint(stone))
		// first rule
		if stone == 0 {
			output = append(output, 1)
		} else if lengthStone%2 == 0 {
			divisor := int64(math.Pow10(lengthStone / 2))
			firstStone := stone / divisor
			secondStone := stone - (firstStone * divisor)
			output = append(output, firstStone, secondStone)
		} else {
			output = append(output, stone*2024)
		}
	}
	fmt.Println(output)
	return output
}
