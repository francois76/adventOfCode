package main

import (
	"bufio"
	"fmt"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
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

	shared.Open("../2/2.txt", func(fileScanner *bufio.Scanner) {
		maxCount += mapResult[fileScanner.Text()]
	})

	fmt.Println(maxCount)
}
