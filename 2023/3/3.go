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
	isCounted  bool
}

func (n *node) hasSymbolAround(position string) bool {
	axis := strings.Split(position, "-")
	x, _ := strconv.ParseInt(axis[0], 10, 64)
	y, _ := strconv.ParseInt(axis[1], 10, 64)
	if val, ok := allElements[fmt.Sprintf("%d-%d", x-1, y)]; ok && val.isSymbol {
		return true
	}
	if val, ok := allElements[fmt.Sprintf("%d-%d", x+1, y)]; ok && val.isSymbol {
		return true
	}
	if val, ok := allElements[fmt.Sprintf("%d-%d", x, y-1)]; ok && val.isSymbol {
		return true
	}
	if val, ok := allElements[fmt.Sprintf("%d-%d", x, y+1)]; ok && val.isSymbol {
		return true
	}
	if val, ok := allElements[fmt.Sprintf("%d-%d", x-1, y-1)]; ok && val.isSymbol {
		return true
	}
	if val, ok := allElements[fmt.Sprintf("%d-%d", x-1, y+1)]; ok && val.isSymbol {
		return true
	}
	if val, ok := allElements[fmt.Sprintf("%d-%d", x+1, y-1)]; ok && val.isSymbol {
		return true
	}
	if val, ok := allElements[fmt.Sprintf("%d-%d", x+1, y+1)]; ok && val.isSymbol {
		return true
	}
	// on est sur le même nombre, du coup on appelle en récursivité l'élément suivant
	if val, ok := allElements[fmt.Sprintf("%d-%d", x, y+1)]; ok && val == n && val.hasSymbolAround(fmt.Sprintf("%d-%d", x, y+1)) {
		return true
	}
	return false
}

func (n *node) value(key string) int64 {
	if n.isCounted {
		return 0
	}
	if n.hasSymbolAround(key) {
		val, _ := strconv.ParseInt(n.innerValue, 10, 64)
		n.isCounted = true
		return val
	}
	return 0
}

func main() {

	shared.Run(func() any {
		total := int64(0)
		currentLine := 0
		var previousElement *node = nil

		shared.Open("3.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			previousElement = nil
			elementsOnCurrentLine := strings.Split(currentLineString, "")
			for idx, element := range elementsOnCurrentLine {

				if element == "." {
					previousElement = nil
					continue
				}
				currentKey := fmt.Sprintf("%d-%d", currentLine, idx)
				if _, err := strconv.ParseInt(element, 10, 64); err != nil {
					previousElement = nil
					// is a symbol, we add it to the current line symbol
					allElements[currentKey] = &node{
						isSymbol: true,
					}
				} else if previousElement != nil {
					previousElement.innerValue = previousElement.innerValue + element
					allElements[currentKey] = previousElement
				} else {
					newElement := &node{
						innerValue: element,
					}
					allElements[currentKey] = newElement
					previousElement = newElement
				}
			}
			currentLine++
		})
		for key, element := range allElements {
			if !element.isSymbol {
				total = total + element.value(key)
			}
		}
		return total
	})

}
