package main

import (
	"bufio"
	"fmt"
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

func (t *tree) leftCriteriaDistance(h int64) int64 {
	if t.left == nil {
		return 0
	}
	if t.left.height >= h {
		return 1
	}
	return 1 + t.left.leftCriteriaDistance(h)
}

func (t *tree) rightCriteriaDistance(h int64) int64 {
	if t.right == nil {
		return 0
	}
	if t.right.height >= h {
		return 1
	}
	return 1 + t.right.rightCriteriaDistance(h)
}

func (t *tree) topCriteriaDistance(h int64) int64 {
	if t.top == nil {
		return 0
	}
	if t.top.height >= h {
		return 1
	}
	return 1 + t.top.topCriteriaDistance(h)
}

func (t *tree) bottomCriteriaDistance(h int64) int64 {
	if t.bottom == nil {
		return 0
	}
	if t.bottom.height >= h {
		return 1
	}
	return 1 + t.bottom.bottomCriteriaDistance(h)
}

func (t *tree) scenicScore() int64 {
	return t.leftCriteriaDistance(t.height) * t.rightCriteriaDistance(t.height) * t.topCriteriaDistance(t.height) * t.bottomCriteriaDistance(t.height)
}

func main() {
	shared.Run(func() any {
		grid := map[int]map[int]*tree{}
		grid[-1] = map[int]*tree{}
		row := 0
		shared.Open("../8/8.txt", func(fileScanner *bufio.Scanner) {
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
		max := int64(0)
		for idx := 0; idx < len(grid); idx++ {
			row := grid[idx]
			fmt.Println()
			for jdx := 0; jdx < len(row); jdx++ {
				tree := row[jdx]
				score := tree.scenicScore()
				if score > max {
					max = score
				}
			}
		}
		return max
	})
}
