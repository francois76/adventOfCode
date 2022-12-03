package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("3.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	// variables
	m := make(map[string]int)

	// Populate the map with keys and values
	for i, ch := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		m[string(ch)+string(ch)] = i + 1
	}

	totalCount := 0

	// loop
	for fileScanner.Scan() {
		totalCount += getCommonChar(m, fileScanner.Text())
	}
	fmt.Println(totalCount)
	readFile.Close()
}

func getCommonChar(m map[string]int, line string) int {
	firstPart := line[:len(line)/2]
	secondpart := line[len(line)/2:]
	for _, x := range firstPart {
		for _, y := range secondpart {
			if rs := m[string(x)+string(y)]; rs != 0 {
				return rs
			}
		}
	}
	return 0
}
