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
	shared.Run(func() interface{} {
		shared.Open("9.txt", func(fileScanner *bufio.Scanner) {
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
			break
		}
		sum += int64(key) * parseInt(value)
	}
	return sum
}

func sort[T any](liste []T, criteria func(item T) bool) {
	n := len(liste)
	right := n - 1
	left := 0

	// Tant qu'on n'a pas parcouru toute la liste
	for left < right {
		// Chercher le premier point depuis la left
		for left < n && !criteria(liste[left]) {
			left++
		}

		// Chercher le premier non-point depuis la right
		for right >= 0 && criteria(liste[right]) {
			right--
		}

		// Si on a trouvé un point à left et un non-point à right
		if left < right {
			// Permuter les éléments
			liste[left], liste[right] = liste[right], liste[left]
		}
	}
}
