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
		shared.Open("../4/4.txt", func(fileScanner *bufio.Scanner) {
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

type M struct {
	letterStruct
}

type A struct {
	letterStruct
}

func (a A) validWordFromHere() int64 {
	leftUp := a.getNeighbor(-1, -1)
	leftDown := a.getNeighbor(-1, 1)
	rightUp := a.getNeighbor(1, -1)
	rightDown := a.getNeighbor(1, 1)
	if leftUp != nil && leftDown != nil && rightUp != nil && rightDown != nil {
		_, leftUpIsM := (*leftUp).(M)
		_, leftDownIsM := (*leftDown).(M)
		_, rightUpIsM := (*rightUp).(M)
		_, rightDownIsM := (*rightDown).(M)
		_, leftUpIsS := (*leftUp).(S)
		_, leftDownIsS := (*leftDown).(S)
		_, rightUpIsS := (*rightUp).(S)
		_, rightDownIsS := (*rightDown).(S)
		if (leftUpIsM && rightDownIsS || leftUpIsS && rightDownIsM) && (leftDownIsM && rightUpIsS || leftDownIsS && rightUpIsM) {
			return 1
		}
	}
	return 0
}

type S struct {
	letterStruct
}
