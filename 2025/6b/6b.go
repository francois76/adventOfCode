package main

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

// Part 2: Cephalopod math = read right-to-left, each column is one number (MSB top).
// Problems separated by all-space columns; operator at bottom row of problem.
func main() {
	shared.Run(func() any {
		lines := []string{}
		shared.Open("../6/6.txt", func(fileScanner *bufio.Scanner) {
			lines = append(lines, fileScanner.Text())
		})
		if len(lines) == 0 {
			return int64(0)
		}

		// Pad lines to same width
		maxW := 0
		for _, l := range lines {
			if len(l) > maxW {
				maxW = len(l)
			}
		}
		grid := make([][]rune, len(lines))
		for i, l := range lines {
			row := make([]rune, maxW)
			for j := 0; j < maxW; j++ {
				if j < len(l) {
					row[j] = rune(l[j])
				} else {
					row[j] = ' '
				}
			}
			grid[i] = row
		}

		lastRow := len(grid) - 1
		total := int64(0)

		// Scan columns right-to-left
		col := maxW - 1
		for col >= 0 {
			// Skip separator columns (all spaces)
			if isEmptyColumn(grid, col) {
				col--
				continue
			}

			// Found start of a problem; consume consecutive non-empty columns
			problemEnd := col
			problemStart := col
			for problemStart > 0 && !isEmptyColumn(grid, problemStart-1) {
				problemStart--
			}

			// Extract operator (bottom row, somewhere in [problemStart..problemEnd])
			op := rune(0)
			for c := problemEnd; c >= problemStart; c-- {
				ch := grid[lastRow][c]
				if ch == '+' || ch == '*' {
					op = ch
					break
				}
			}

			// Parse numbers: ALL columns (including operator column) form numbers from rows 0 to lastRow-1
			// The operator is only at lastRow, so we read digits above it
			values := []int64{}
			for c := problemEnd; c >= problemStart; c-- {
				// Build number from top to lastRow-1 (above the operator row)
				numStr := ""
				for r := 0; r < lastRow; r++ {
					ch := grid[r][c]
					if ch >= '0' && ch <= '9' {
						numStr += string(ch)
					}
				}
				if numStr != "" {
					n, _ := strconv.ParseInt(strings.TrimSpace(numStr), 10, 64)
					values = append(values, n)
				}
			}

			// Apply operator
			result := int64(0)
			switch op {
			case '+':
				for _, v := range values {
					result += v
				}
			case '*':
				result = 1
				for _, v := range values {
					result *= v
				}
			}
			total += result

			col = problemStart - 1
		}

		return total
	})
}

func isEmptyColumn(grid [][]rune, col int) bool {
	for r := 0; r < len(grid); r++ {
		if grid[r][col] != ' ' {
			return false
		}
	}
	return true
}
