package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("../2/2.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	maxCount := 0

	// X for lose, Y for draw, and Z for win
	// A for Rock, B for Paper, and C for Scissors

	// shape points
	// 1 for rock,
	// 2 for paper
	// 3 for scisors

	// 0 if you lost, 3 if the round was a draw, and 6 if you won

	mapResult := map[string]int{
		"A X": 0 + 3, // lose + scisors
		"A Y": 3 + 1, // draw + rock
		"A Z": 6 + 2, // win + paper
		"B X": 0 + 1, // lose + rock
		"B Y": 3 + 2, // draw + paper
		"B Z": 6 + 3, // win + scisors
		"C X": 0 + 2, // lose + paper
		"C Y": 3 + 3, // draw + scisors
		"C Z": 6 + 1, // win + rock
	}

	for fileScanner.Scan() {
		maxCount += mapResult[fileScanner.Text()]
	}
	fmt.Println(maxCount)
	readFile.Close()
}
