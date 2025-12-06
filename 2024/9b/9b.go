package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	items := []string{}
	shared.Run(func() any {
		shared.Open("../9/9.txt", func(fileScanner *bufio.Scanner) {
			line := fileScanner.Text()
			index := 0
			empty := false
			for _, letter := range strings.Split(line, "") {
				for range parseInt(letter) {
					if empty {
						items = append(items, ".")
					} else {
						items = append(items, fmt.Sprint(index))
					}

				}
				if !empty {
					index++
				}
				empty = !empty
			}

		})
		fmt.Println(items)
		sort(items, func(item string) bool {
			return item == "."
		})
		fmt.Println(items)
		return checksum(items)
	})
}
func parseInt(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func checksum(items []string) int64 {
	sum := int64(0)
	for key, value := range items {
		if value == "." {
			continue
		}
		sum += int64(key) * parseInt(value)
	}
	return sum
}

func sort[T comparable](list []T, criteria func(item T) bool) {
	n := len(list)
	right := n - 1

	// we are progressing both left and right but right is
	// only decreasing
	for 0 < right {
		// time.Sleep(1 * time.Second)
		left := 0
		spaceToMove := 0

		// searching first non empty slot from right
		for right >= 0 && criteria(list[right]) {
			right--
		}

		for list[right-spaceToMove] == list[right] {
			spaceToMove++
			if right-spaceToMove < 0 {
				// right has reached the left border
				break
			}
		}

		// searching first empty slot from left
		for {
			if left < n && !criteria(list[left]) {
				left++
			} else {
				// an empty value has been found. Checking if there is enough space
				// if we havent enough space, loop isn't stopped contrary to day 9.
				hasEnoughSpace := true
				for idx := range spaceToMove {
					if len(list) <= left+idx {
						// it means that we got completely right and we cant find a slot big enough
						// so we wont move that item and will reduce directly right by the
						// offset
						break
					}
					if !criteria(list[left+idx]) {
						hasEnoughSpace = false
						// to go quicker, we directly increase the left to the size of the slot we
						// just checked
						left += idx
						break
					}
				}
				if hasEnoughSpace {
					break
				}
			}
		}

		// if left item is still before right item, we swap
		if left < right {
			for idx := range spaceToMove {
				// swapping items
				list[left+idx], list[right-idx] = list[right-idx], list[left+idx]
			}
		} else {
			// we found a new place but its more on right that current position. Decreasing right and
			// doing nothing
			right -= spaceToMove
		}
	}
}
