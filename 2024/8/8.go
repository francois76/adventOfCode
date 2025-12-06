package main

import (
	"bufio"
	"fmt"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	shared.Run(func() any {
		idx := 0
		grid := grid{}
		nodesByFrequency := map[string][]*Node{}
		shared.Open("8.txt", func(fileScanner *bufio.Scanner) {
			line := fileScanner.Text()
			grid = append(grid, []Element{})
			for jdx, c := range line {
				switch c {
				case '.':
					grid[len(grid)-1] = append(grid[len(grid)-1], Empty{})
				default:
					frequency := string(c)
					if _, ok := nodesByFrequency[frequency]; !ok {
						nodesByFrequency[frequency] = []*Node{}
					}
					node := &Node{grid: &grid, x: idx, y: jdx, frequency: frequency}
					nodesByFrequency[frequency] = append(nodesByFrequency[frequency], node)
					grid[len(grid)-1] = append(grid[len(grid)-1], node)
				}
			}
			idx++
		})
		antinodeCount := 0
		for _, nodes := range nodesByFrequency {
			for _, n1 := range nodes {
				for _, n2 := range nodes {
					if n1 == n2 {
						continue
					}
					vectorX, vectorY := n1.Vector(*n2)
					if grid.AddAntinode(n2.x+vectorX, n2.y+vectorY) {
						antinodeCount++
					}
				}
			}
		}
		fmt.Println(grid)
		return antinodeCount
	})
}

type grid [][]Element

func (g grid) GetElement(x int, y int) Element {
	return g[x][y]
}

func (g grid) IsOutOfBounds(x int, y int) bool {
	return x < 0 || x >= len(g) || y < 0 || y >= len(g[0])
}

func (g grid) AddAntinode(x int, y int) bool {
	if g.IsOutOfBounds(x, y) {
		fmt.Println("skipping ", x, " ", y, " because out of bound")
		return false
	}
	destination := g[x][y]
	if ok := destination.CanBeNewAntinode(); !ok {
		fmt.Println("skipping ", x, " ", y, " because place is taken")
		return false
	}

	g[x][y] = Antinode{}
	return true
}

func (g grid) String() string {
	fmt.Print("\033[H\033[2J") // clean screen
	output := ""
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			output += g[i][j].String()
		}
		output += "\n"
	}
	return output
}

type Element interface {
	String() string
	CanBeNewAntinode() bool
}

type Empty struct {
}

func (e Empty) String() string {
	return "."
}

func (e Empty) CanBeNewAntinode() bool {
	return true
}

type Antinode struct {
}

func (e Antinode) String() string {
	return "#"
}

func (e Antinode) CanBeNewAntinode() bool {
	return false
}

type Node struct {
	frequency string
	x         int
	y         int
	grid      *grid
}

func (n1 Node) Vector(n2 Node) (int, int) {
	return n2.x - n1.x, n2.y - n1.y
}

func (e Node) String() string {
	return e.frequency
}

func (e Node) CanBeNewAntinode() bool {
	return true
}
