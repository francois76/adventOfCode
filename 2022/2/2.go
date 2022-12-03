package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("2.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	maxCount := 0

	// X for Rock, Y for Paper, and Z for Scissors
	// A for Rock, B for Paper, and C for Scissors

	// shape points
	// 1 for X,
	// 2 for Y
	// 3 for Z

	// 0 if you lost, 3 if the round was a draw, and 6 if you won

	mapResult := map[string]int{
		"A X": 3 + 1, // draw + 1
		"A Y": 6 + 2, // win + 2
		"A Z": 0 + 3, // lose +3
		"B X": 0 + 1, // lose +1
		"B Y": 3 + 2, // draw +2
		"B Z": 6 + 3, // win +3
		"C X": 6 + 1, // win +1
		"C Y": 0 + 2, // lose +2
		"C Z": 3 + 3, // draw +3
	}

	for fileScanner.Scan() {
		maxCount += mapResult[fileScanner.Text()]
	}
	fmt.Println(maxCount)
	readFile.Close()
}
