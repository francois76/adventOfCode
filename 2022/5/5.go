package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	shared.Run(func() interface{} {
		stacks := map[int64][]string{}
		shared.Open("5_init.txt", func(fileScanner *bufio.Scanner) {
			rawLine := fileScanner.Text()
			line := strings.ReplaceAll(rawLine, "]    ", "] @")
			line = strings.ReplaceAll(line, "    [", "@ [")
			line = strings.ReplaceAll(line, "    ", "@ ")
			lineSplit := strings.Split(line, " ")
			for idx, item := range lineSplit {
				if item == "@" {
					continue
				}
				if item == "1" {
					break
				}
				stacks[int64(idx+1)] = append([]string{item}, stacks[int64(idx+1)]...)
			}
		})
		fmt.Println(fmt.Sprintf("%v", stacks))
		shared.Open("5_steps.txt", func(fileScanner *bufio.Scanner) {
			lineSplit := strings.Split(fileScanner.Text(), " ")
			number, _ := strconv.ParseInt(lineSplit[1], 10, 64)
			from, _ := strconv.ParseInt(lineSplit[3], 10, 64)
			to, _ := strconv.ParseInt(lineSplit[5], 10, 64)
			payload := stacks[from][len(stacks[from])-int(number):]
			stacks[from] = stacks[from][:len(stacks[from])-int(number)]
			fmt.Println(number)
			for idx := number - 1; idx >= 0; idx-- {
				stacks[to] = append(stacks[to], payload[idx])
			}
		})
		fmt.Println(fmt.Sprintf("%v", stacks))
		fmt.Println()
		for idx := int64(1); idx <= int64(9); idx++ {
			fmt.Print(string(stacks[idx][len(stacks[idx])-1][1]))
		}
		fmt.Println()
		return ""
	})
}
