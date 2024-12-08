package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
	"github.com/samber/lo"
)

func main() {
	scoresToBid := map[int64]int64{}
	scoreToHand := map[int64]string{}
	result := int64(0)
	shared.Run(func() interface{} {
		shared.Open("7.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			elementsOnCurrentLine := strings.Split(currentLineString, " ")
			hand := elementsOnCurrentLine[0]
			bid, _ := strconv.ParseInt(elementsOnCurrentLine[1], 10, 64)
			score := computeScore(hand)
			multiplier := score * 10
			// check if another exact store exist
			if h, ok := scoreToHand[multiplier]; ok {
				if h < hand {
					multiplier = score*10 + 1
				} else {
					multiplier = score*10 - 1
				}
			}
			scoreToHand[multiplier] = hand
			scoresToBid[multiplier] = bid
		})
		rank := int64(1)
		fmt.Println("max: ", int64(14*Flush)+1)
		fmt.Println(fmt.Printf("%v", scoreToHand))
		for i := range int64(14*Flush)*10 + 1 {
			if value, ok := scoresToBid[i]; ok {
				result += value * rank
				fmt.Println("hand:  ", scoreToHand[i], " ", scoresToBid[i])
				rank++
			}
		}
		return result
	})
}

type RankMultiplier int64

const (
	Flush       RankMultiplier = 15 * Square
	Square      RankMultiplier = 15 * Brelan
	Brelan      RankMultiplier = 15 * DoublePair
	DoublePair  RankMultiplier = 15 * SinglePair
	SinglePair  RankMultiplier = 15 * HighestCard
	HighestCard RankMultiplier = 1
)

var letterToValue = map[string]int64{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

// computeScore takes a poker hand and determining a score so that a better hand gives a better score
func computeScore(hand string) int64 {
	type cardNumber struct {
		cardName string
		number   int
	}
	items := []cardNumber{}
	for _, letter := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"} {
		items = append(items, cardNumber{letter, strings.Count(hand, letter)})
	}

	findN := func(n int, items []cardNumber) (cardNumber, bool) {
		return lo.Find(items, func(item cardNumber) bool {
			return item.number == n
		})
	}
	findHighestOne := func(items []cardNumber) int64 {
		return lo.Max(lo.Map(lo.Filter(items, func(item cardNumber, index int) bool {
			return item.number == 1
		}), func(item cardNumber, index int) int64 {
			return letterToValue[item.cardName]
		}))
	}
	// flush
	if item, ok := findN(5, items); ok {
		return letterToValue[item.cardName] * int64(Flush)
	}

	// square
	if item, ok := findN(4, items); ok {
		complement, _ := findN(1, items)
		return letterToValue[item.cardName]*int64(Square) + letterToValue[complement.cardName]
	}

	// full and brelan
	if item, ok := findN(3, items); ok {
		// full
		if complement, ok := findN(2, items); ok {
			return letterToValue[item.cardName]*int64(Brelan) + letterToValue[complement.cardName]*int64(DoublePair)
		}
		// brelan
		return letterToValue[item.cardName]*int64(Brelan) + findHighestOne(items)
	}

	// two pairs
	if items := lo.Filter(items, func(item cardNumber, index int) bool {
		return item.number == 2
	}); len(items) == 2 {
		value1 := letterToValue[items[0].cardName]
		value2 := letterToValue[items[1].cardName]
		value3 := findHighestOne(items)
		if value1 > value2 {
			return value1*int64(DoublePair) + value2*int64(SinglePair) + value3
		}
		return value2*int64(DoublePair) + value1*int64(SinglePair) + value3
	}
	// one pair
	if item, ok := findN(2, items); ok {
		return letterToValue[item.cardName]*int64(SinglePair) + findHighestOne(items)
	}
	return findHighestOne(items)
}
