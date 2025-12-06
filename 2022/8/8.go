package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

type tree struct {
	height int64
	left   *tree
	right  *tree
	top    *tree
	bottom *tree
}

func (t *tree) leftCriteriaVisibility() int64 {
	if t.left == nil {
		return -1
	}
	return int64(math.Max(float64(t.left.leftCriteriaVisibility()), float64(t.left.height)))
}

func (t *tree) rightCriteriaVisibility() int64 {
	if t.right == nil {
		return -1
	}
	return int64(math.Max(float64(t.right.rightCriteriaVisibility()), float64(t.right.height)))
}

func (t *tree) topCriteriaVisibility() int64 {
	if t.top == nil {
		return -1
	}
	return int64(math.Max(float64(t.top.topCriteriaVisibility()), float64(t.top.height)))
}

func (t *tree) bottomCriteriaVisibility() int64 {
	if t.bottom == nil {
		return -1
	}
	return int64(math.Max(float64(t.bottom.bottomCriteriaVisibility()), float64(t.bottom.height)))
}

func (t *tree) isVisible() bool {
	if t.left == nil || t.right == nil || t.top == nil || t.bottom == nil {
		return true
	}
	return t.leftCriteriaVisibility() < t.height || t.rightCriteriaVisibility() < t.height || t.topCriteriaVisibility() < t.height || t.bottomCriteriaVisibility() < t.height
}

func main() {
	shared.Run(func() any {
		grid := map[int]map[int]*tree{}
		grid[-1] = map[int]*tree{}
		row := 0
		shared.Open("8.txt", func(fileScanner *bufio.Scanner) {
			line := strings.Split(fileScanner.Text(), "")
			grid[row] = map[int]*tree{}
			for col := 0; col < len(line); col++ {
				value := line[col]
				height, _ := strconv.ParseInt(value, 10, 64)
				tree := tree{
					height: height,
					left:   grid[row][col-1],
					top:    grid[row-1][col],
				}
				grid[row][col] = &tree
				if grid[row][col-1] != nil {
					grid[row][col-1].right = &tree
				}
				if grid[row-1][col] != nil {
					grid[row-1][col].bottom = &tree
				}
			}
			row++
		})
		counter := 0
		for idx := 0; idx < len(grid); idx++ {
			row := grid[idx]
			fmt.Println()
			for jdx := 0; jdx < len(row); jdx++ {
				tree := row[jdx]
				if tree.isVisible() {
					fmt.Print("X")
					counter++
				} else {
					fmt.Print("_")
				}
			}
		}
		return counter
	})
}
