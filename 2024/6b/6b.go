package main

import (
	"bufio"
	"errors"
	"fmt"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	shared.Run(func() interface{} {
		grid := grid{}
		var player *Player
		idx := 0
		shared.Open("../6/6.txt", func(fileScanner *bufio.Scanner) {
			line := fileScanner.Text()
			grid = append(grid, []Element{})

			for jdx, c := range line {
				switch c {
				case '#':
					grid[len(grid)-1] = append(grid[len(grid)-1], Obstacle{})
				case '.':
					grid[len(grid)-1] = append(grid[len(grid)-1], Vide{})
				case '^':
					player = &Player{orientation: Up, grid: &grid, x: idx, y: jdx}
					grid[len(grid)-1] = append(grid[len(grid)-1], player)
				}
			}
			idx++
		})
		blockers := 0
		for {
			nextX, nextY := player.NextPosition()
			if grid.IsOutOfBounds(nextX, nextY) {
				player.traceLength++
				break
			}
			if !grid.canTrace(nextX, nextY, (player.orientation+1)%4) {
				blockers++
				fmt.Println(nextX, " ", nextY, " ")
				// time.Sleep(3 * time.Second)
			}
			if err := player.grid.MovePlayerTo(player, nextX, nextY); err != nil {
				player.TurnRight()
			}
			// fmt.Println(grid)
			// time.Sleep(time.Second / 10)
		}

		return blockers
	})

}

type Player struct {
	x           int
	y           int
	traceLength int
	orientation orientation
	grid        *grid
}

func (p *Player) TurnRight() {
	p.orientation = (p.orientation + 1) % 4
}

func (p *Player) NextPosition() (int, int) {
	return NextPosition(p.orientation, p.x, p.y)
}

func NextPosition(orientation orientation, x int, y int) (int, int) {
	switch orientation {
	case Up:
		return x - 1, y
	case Right:
		return x, y + 1
	case Down:
		return x + 1, y
	case Left:
		return x, y - 1
	}
	return 0, 0
}

func (p Player) String() string {
	switch p.orientation {
	case Up:
		return "^"
	case Right:
		return ">"
	case Down:
		return "v"
	case Left:
		return "<"
	}
	return ""
}

type grid [][]Element

func (g grid) GetElement(x int, y int) Element {
	return g[x][y]
}

func (g grid) MovePlayerTo(player *Player, nextX int, nextY int) error {
	destination := g[nextX][nextY]
	if _, ok := destination.(Obstacle); ok {
		return errors.New("obstacle")
	}

	// when we move, we bump the counter if there werent already a trace
	if _, ok := destination.(Trace); !ok {
		player.traceLength++
	}
	g[player.x][player.y] = Trace{
		orientation: player.orientation,
	}
	g[nextX][nextY] = player
	player.x = nextX
	player.y = nextY
	return nil
}

func (g grid) canTrace(x int, y int, orientation orientation) bool {
	paths := map[string]struct{}{}
	for {
		pathKey := fmt.Sprintf("%d-%d-%d", x, y, orientation)
		// permit to detect cycles that dont correspond with traced patterns
		if _, ok := paths[pathKey]; ok {
			return false
		}
		paths[pathKey] = struct{}{}
		// getting element of current position
		e := g.GetElement(x, y)
		t, ok := e.(Trace)
		x, y = NextPosition(orientation, x, y)
		// we reached the bound so it means we can trace
		if g.IsOutOfBounds(x, y) {
			return true
		}
		e = g.GetElement(x, y)
		if _, ok := e.(Obstacle); ok {
			orientation = (orientation + 1) % 4
		}
		// checking the current trace with the new orientation that is the correct one
		if ok && t.orientation == orientation {
			return false
		}
	}
}

func (g grid) IsOutOfBounds(x int, y int) bool {
	return x < 0 || x >= len(g) || y < 0 || y >= len(g[0])
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
}

type Obstacle struct {
}

func (e Obstacle) String() string {
	return "#"
}

type Vide struct {
}

func (e Vide) String() string {
	return "."
}

type Trace struct {
	orientation orientation
}

func (e Trace) String() string {
	switch e.orientation {
	case Up:
		return "U"
	case Down:
		return "D"
	case Left:
		return "L"
	case Right:
		return "R"
	default:
		return "X"
	}
}

type orientation = int

const (
	Up    orientation = iota
	Right orientation = iota
	Down  orientation = iota
	Left  orientation = iota
)
