package main

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
	"github.com/samber/lo"
)

func main() {
	shared.Run(func() any {

		total := int64(0)
		shared.Open("../7/7.txt", func(fileScanner *bufio.Scanner) {
			currentLineParts := strings.Split(fileScanner.Text(), ": ")
			totalToFind := parseInt(currentLineParts[0])
			items := lo.Map(strings.Split(currentLineParts[1], " "), func(item string, _ int) int64 {
				return parseInt(item)
			})
			total += ComputeTotal(totalToFind, items)
		})
		return total
	})
}

func ComputeTotal(totalToFind int64, items []int64) int64 {
	layer := []int64{}
	for idx, item := range items {
		if idx == 0 {
			layer = append(layer, item)
		} else {
			layer = nextLayer(layer, item)
		}
	}
	if slices.Contains(layer, totalToFind) {
		return totalToFind
	}
	return 0
}

func nextLayer(layer []int64, value int64) []int64 {
	result := []int64{}
	for _, item := range layer {
		result = append(result, item+value)
		result = append(result, item*value)
		result = append(result, concatenate(item, value))
	}
	return result
}

func parseInt(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func concatenate(a, b int64) int64 {
	return parseInt(fmt.Sprintf("%d%d", a, b))
}
