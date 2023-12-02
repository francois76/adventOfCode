package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

type position struct {
	x     int64
	y     int64
	print bool
}

func (p *position) distance(p2 position) int64 {
	fmt.Println("(", p2.x, p.x, ")(", p2.y, p.y, ")")
	return (p2.x - p.x) * (p2.x - p.x)
}

func (p *position) String() string {
	return fmt.Sprint(p.x, "-", p.y)
}

func main() {
	tail := position{0, 0, true}
	head := position{0, 0, true}
	shared.Run(func() interface{} {
		positions := map[string]bool{}

		shared.Open("9.txt", func(fileScanner *bufio.Scanner) {
			elements := strings.Split(fileScanner.Text(), " ")
			iterations, _ := strconv.ParseInt(elements[1], 10, 64)
			var f func(*position, bool)
			switch elements[0] {
			case "U":
				f = func(p *position, print bool) {
					p.x++
					p.print = print
				}
			case "D":
				f = func(p *position, print bool) {
					p.x--
					p.print = print
				}
			case "L":
				f = func(p *position, print bool) {
					p.y--
					p.print = print
				}
			case "R":
				f = func(p *position, print bool) {
					p.y++
					p.print = print
				}
			}
			for idx := int64(0); idx < iterations; idx++ {
				if tail.print {
					positions[tail.String()] = true
				}

				for i := 4; i >= 0; i-- {
					fmt.Println()
					for j := 0; j < 5; j++ {
						if positions[fmt.Sprint(i, "-", j)] {
							fmt.Print("X")
						} else {
							fmt.Print("_")
						}
					}
				}
				fmt.Println()
				head.distance(tail)
				f(&head, true)
				fmt.Println()
				d := head.distance(tail) == 0
				f(&tail, d)
				fmt.Println()
				head.distance(tail)

				fmt.Println()
				fmt.Println()
				fmt.Println()
			}

		})
		return len(positions)
	})
}
