package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

const MAX_LINE_COUNT = 5

func main() {
	shared.Run(func() any {
		count := int64(0)
		lineCount := 1
		items := [][]int64{}
		shared.Open("6.txt", func(fileScanner *bufio.Scanner) {
			line := fileScanner.Text()
			if lineCount > MAX_LINE_COUNT {
				return
			}
			if lineCount == MAX_LINE_COUNT {
				fmt.Println(items)
				itemLine := strings.Split(line, " ")
				index := 0
				for _, item := range itemLine {
					switch item {
					case "":
						continue
					case "*":
						fmt.Println("multiplying")
						subCount := int64(1)
						for _, itemInts := range items {
							fmt.Print(itemInts[index], ", ")
							subCount *= itemInts[index]
						}
						count += subCount
						fmt.Println("count is temporary at ", count)
						index++
					case "+":
						fmt.Println("adding")
						subCount := int64(0)
						for _, itemInts := range items {
							fmt.Print(itemInts[index], ", ")
							subCount += itemInts[index]
						}
						count += subCount
						fmt.Println("count is temporary at ", count)
						index++
					}
				}
				return
			}
			lineCount++
			itemLine := strings.Split(line, " ")
			itemInts := []int64{}
			for _, item := range itemLine {
				if item != "" {
					i, _ := strconv.ParseInt(item, 10, 64)
					itemInts = append(itemInts, i)
				}
			}
			items = append(items, itemInts)
		})
		return count
	})
}
