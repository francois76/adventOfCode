package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
	"github.com/samber/lo"
)

var allElements = map[string]*node{}

type node struct {
	innerValue string
	isGear     bool
}

func (n *node) gearProduct(position string) int64 {
	axis := strings.Split(position, "-")
	x, _ := strconv.ParseInt(axis[0], 10, 64)
	y, _ := strconv.ParseInt(axis[1], 10, 64)
	referencesAround := map[*node]bool{}
	referencesAround[allElements[fmt.Sprintf("%d-%d", x-1, y)]] = true
	referencesAround[allElements[fmt.Sprintf("%d-%d", x+1, y)]] = true
	referencesAround[allElements[fmt.Sprintf("%d-%d", x, y-1)]] = true
	referencesAround[allElements[fmt.Sprintf("%d-%d", x, y+1)]] = true
	referencesAround[allElements[fmt.Sprintf("%d-%d", x-1, y-1)]] = true
	referencesAround[allElements[fmt.Sprintf("%d-%d", x-1, y+1)]] = true
	referencesAround[allElements[fmt.Sprintf("%d-%d", x+1, y-1)]] = true
	referencesAround[allElements[fmt.Sprintf("%d-%d", x+1, y+1)]] = true
	delete(referencesAround, nil)
	if len(referencesAround) == 2 {
		keys := lo.Keys(referencesAround)
		return keys[0].value() * keys[1].value()
	}
	fmt.Println("ignoring ", lo.Keys(referencesAround)[0].value())
	return 0
}

func (n *node) value() int64 {
	val, _ := strconv.ParseInt(n.innerValue, 10, 64)
	return val
}

func main() {

	shared.Run(func() interface{} {
		total := int64(0)
		currentLine := 0
		var previousElement *node = nil

		shared.Open("../3/3.txt", func(fileScanner *bufio.Scanner) {
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
					if element == "*" {
						allElements[currentKey] = &node{
							isGear: true,
						}
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
			if element.isGear {
				total = total + element.gearProduct(key)
			}
		}
		return total
	})

}
