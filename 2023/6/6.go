package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("example: ", computeOneItem(7, 9)*computeOneItem(15, 40)*computeOneItem(30, 200))
	fmt.Println("part 1: ", computeOneItem(41, 214)*computeOneItem(96, 1789)*computeOneItem(88, 1127)*computeOneItem(94, 1055))
	fmt.Println("part 2: ", computeOneItem(41968894, 214178911271055))
}

// return the number of possibility for one peer of item.
// actually the problem corresponds to an inequation k(T-K) > D
func computeOneItem(T, D float64) int64 {
	// to handle float values with integer interval, we take floor and ceils that give alway the largest corresponding interval and we just have to process max-min-1
	return int64(math.Ceil((T+math.Sqrt(T*T-4*D))/2)-math.Floor((T-math.Sqrt(T*T-4*D))/2)) - 1
}
