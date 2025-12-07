package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

var allElements = map[string]*node{}

type node struct {
	innerValue string
	isSymbol   bool
}

func (n *node) isAccessible(position string) bool {
	if !n.isSymbol {
		return false
	}
	axis := strings.Split(position, "-")
	x, _ := strconv.ParseInt(axis[0], 10, 64)
	y, _ := strconv.ParseInt(axis[1], 10, 64)
	count := 0
	if val, ok := allElements[fmt.Sprintf("%d-%d", x-1, y)]; ok && val.isSymbol {
		count++
	}
	if val, ok := allElements[fmt.Sprintf("%d-%d", x+1, y)]; ok && val.isSymbol {
		count++
	}
	if val, ok := allElements[fmt.Sprintf("%d-%d", x, y-1)]; ok && val.isSymbol {
		count++
	}
	if val, ok := allElements[fmt.Sprintf("%d-%d", x, y+1)]; ok && val.isSymbol {
		count++
	}
	if val, ok := allElements[fmt.Sprintf("%d-%d", x-1, y-1)]; ok && val.isSymbol {
		count++
	}
	if val, ok := allElements[fmt.Sprintf("%d-%d", x-1, y+1)]; ok && val.isSymbol {
		count++
	}
	if val, ok := allElements[fmt.Sprintf("%d-%d", x+1, y-1)]; ok && val.isSymbol {
		count++
	}
	if val, ok := allElements[fmt.Sprintf("%d-%d", x+1, y+1)]; ok && val.isSymbol {
		count++
	}
	if count < 4 {
		n.isSymbol = false
		return true
	}

	return false
}

func main() {

	shared.Run(func() any {
		total := int64(0)
		currentLine := 0

		shared.Open("../4/4.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			elementsOnCurrentLine := strings.Split(currentLineString, "")
			for idx, element := range elementsOnCurrentLine {

				if element == "." {
					continue
				}
				currentKey := fmt.Sprintf("%d-%d", currentLine, idx)
				newElement := &node{
					innerValue: element,
					isSymbol:   element == "@",
				}
				allElements[currentKey] = newElement
			}
			currentLine++
		})

		for {
			for key, element := range allElements {
				if element.isAccessible(key) {
					total++
				}
			}
			fmt.Println(total)
		}

	})

}
