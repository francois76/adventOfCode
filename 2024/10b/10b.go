package main

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/francois76/adventOfCode/shared"
	"github.com/samber/lo"
)

func main() {
	shared.Run(func() interface{} {
		grid := grid{}
		idx := 0
		shared.Open("../10/10.txt", func(fileScanner *bufio.Scanner) {
			line := fileScanner.Text()
			grid = append(grid, []*Element{})

			for jdx, c := range line {
				value := parseInt(string(c))
				summits := []string{}
				if value == 9 {
					//summit can attein itself so its registered immediately
					summits = append(summits, fmt.Sprintf("%d:%d", idx, jdx))
				}
				grid[len(grid)-1] = append(grid[len(grid)-1], &Element{value: int(value), atteignableSubmit: summits})
			}
			idx++
		})
		for idx := 8; idx > 0; idx-- {
			_ = checkLayer(idx, grid)
		}
		return checkLayer(0, grid)
	})
}

func checkLayer(value int, grid grid) int {
	result := 0
	checkNeighbour := func(x, y int) (result []string) {
		result = []string{}
		if grid.IsOutOfBounds(x, y) {
			return
		}
		element := grid.GetElement(x, y)
		if element.value != value+1 {
			return
		}
		result = append(result, lo.Map(element.atteignableSubmit, func(item string, _ int) string {
			return fmt.Sprintf("%d:%d;%s", x, y, item)
		})...)
		return
	}
	for idx, row := range grid {
		for jdx, element := range row {
			if element.value != value {
				continue
			}
			atteignableSummits := append([]string{}, checkNeighbour(idx+1, jdx)...)
			atteignableSummits = append(atteignableSummits, checkNeighbour(idx-1, jdx)...)
			atteignableSummits = append(atteignableSummits, checkNeighbour(idx, jdx+1)...)
			atteignableSummits = append(atteignableSummits, checkNeighbour(idx, jdx-1)...)
			element.atteignableSubmit = atteignableSummits
			result += len(element.atteignableSubmit)
		}
	}
	return result
}

func parseInt(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

type grid [][]*Element

func (g grid) GetElement(x int, y int) *Element {
	return g[x][y]
}
func (g grid) IsOutOfBounds(x int, y int) bool {
	return x < 0 || x >= len(g) || y < 0 || y >= len(g[0])
}

type Element struct {
	value             int
	atteignableSubmit []string
}
