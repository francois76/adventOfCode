package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("../3/3.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	// variables
	m := make(map[string]int)

	// Populate the map with keys and values
	for i, ch := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		m[string(ch)+string(ch)+string(ch)] = i + 1
	}

	totalCount := 0

	// loop
	for fileScanner.Scan() {
		firstPart := fileScanner.Text()
		fileScanner.Scan()
		secondPart := fileScanner.Text()
		fileScanner.Scan()
		totalCount += getCommonChar(m, firstPart, secondPart, fileScanner.Text())

	}
	fmt.Println(totalCount)
	readFile.Close()
}

func getCommonChar(m map[string]int, firstPart string, secondPart string, thirdPart string) int {
	for _, x := range firstPart {
		for _, y := range secondPart {
			for _, z := range thirdPart {
				if rs := m[string(x)+string(y)+string(z)]; rs != 0 {
					return rs
				}
			}
		}
	}
	return 0
}
