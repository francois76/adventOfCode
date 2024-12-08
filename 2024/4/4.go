package main

import (
	"bufio"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

func main() {

	/*

	 O--------------------------------------> Y
	 |
	 |
	 |
	 |
	 |
	 |
	 |
	 |
	 |
	 |
	 X
	*/
	shared.Run(func() interface{} {
		total := int64(0)
		idx := 0
		grid := [][]letter{}
		shared.Open("4.txt", func(fileScanner *bufio.Scanner) {
			grid = append(grid, []letter{})
			for jdx, l := range strings.Split(fileScanner.Text(), "") {
				switch l {
				case "X":
					grid[idx] = append(grid[idx], X{letterStruct{grid: &grid, posX: int64(idx), posY: int64(jdx)}})
				case "M":
					grid[idx] = append(grid[idx], M{letterStruct{grid: &grid, posX: int64(idx), posY: int64(jdx)}})
				case "A":
					grid[idx] = append(grid[idx], A{letterStruct{grid: &grid, posX: int64(idx), posY: int64(jdx)}})
				case "S":
					grid[idx] = append(grid[idx], S{letterStruct{grid: &grid, posX: int64(idx), posY: int64(jdx)}})
				}
			}
			idx++
		})
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				total += grid[i][j].validWordFromHere()
			}
		}
		return total
	})

}

type letter interface {
	validWordFromHere() int64
	neighborIsValid(int64, int64) bool
}

type letterStruct struct {
	posX int64
	posY int64
	grid *[][]letter
}

func (l letterStruct) validWordFromHere() int64 {
	return 0
}

func (l letterStruct) getNeighbor(relX int64, relY int64) *letter {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	n := (*l.grid)[int(l.posX+relX)][int(l.posY+relY)]
	return &n
}

type X struct {
	letterStruct
}

// neighborIsValid implements letter.
func (x X) neighborIsValid(relX int64, relY int64) bool {
	n := x.getNeighbor(relX, relY)
	if n == nil {
		return false
	}
	_, ok := (*n).(M)
	return ok && (*n).neighborIsValid(relX, relY)
}

func (x X) validWordFromHere() int64 {
	trueToInt := func(b bool) int64 {
		if b {
			return 1
		}
		return 0
	}
	numberOfValid := trueToInt(x.neighborIsValid(1, 0)) + trueToInt(x.neighborIsValid(-1, 0)) + trueToInt(x.neighborIsValid(0, 1)) + trueToInt(x.neighborIsValid(0, -1)) +
		trueToInt(x.neighborIsValid(1, 1)) + trueToInt(x.neighborIsValid(-1, 1)) + trueToInt(x.neighborIsValid(1, -1)) + trueToInt(x.neighborIsValid(-1, -1))
	return numberOfValid
}

type M struct {
	letterStruct
}

// neighborIsValid implements letter.
func (m M) neighborIsValid(relX int64, relY int64) bool {
	n := m.getNeighbor(relX, relY)
	if n == nil {
		return false
	}
	_, ok := (*n).(A)
	return ok && (*n).neighborIsValid(relX, relY)
}

type A struct {
	letterStruct
}

// neighborIsValid implements letter.
func (a A) neighborIsValid(relX int64, relY int64) bool {
	n := a.getNeighbor(relX, relY)
	if n == nil {
		return false
	}
	_, ok := (*n).(S)
	return ok && (*n).neighborIsValid(relX, relY)
}

type S struct {
	letterStruct
}

// neighborIsValid implements letter.
func (s S) neighborIsValid(relX int64, relY int64) bool {
	return true
}
