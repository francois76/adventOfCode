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
	// determine for each lenghts of numbers what twin we can do
	twins := map[int64][]int{
		2:  {2},
		3:  {3},
		4:  {2, 4},
		5:  {5},
		6:  {2, 3, 6},
		7:  {7},
		8:  {2, 4, 8},
		9:  {3, 6, 9},
		10: {2, 5, 10},
	}
	buildTwinNumber := func(n int64, twinSize int) int64 {
		args := []any{}
		pattern := ""
		for range twinSize {
			pattern += "%d"
			args = append(args, n)
		}
		result, _ := strconv.ParseInt(fmt.Sprintf(pattern, args...), 10, 64)
		return result
	}
	shared.Run(func() any {
		count := int64(0)

		shared.Open("../2/2.txt", func(fileScanner *bufio.Scanner) {
			currentLineString := fileScanner.Text()
			items := strings.Split(currentLineString, "-")
			start, _ := strconv.ParseInt(items[0], 10, 64)
			allTwins := map[int64]bool{}
			for cut := 0; cut < len(items[0]); cut++ {
				startStart, _ := strconv.ParseInt(items[0][0:cut], 10, 64)
				end, _ := strconv.ParseInt(items[1], 10, 64)
				lenStart := int64(len(items[0]))
				lenEnd := int64(len(items[1]))
				possibleTwins := lo.Union(twins[lenStart], twins[lenEnd])

				for _, twinSize := range possibleTwins {
					startInc := startStart
					for {
						i := buildTwinNumber(startInc, twinSize)
						if i > 1000000000000 {
							break
						}
						if i < start {
							// could happen if begin part smaller than end part
							startInc++
							continue
						}
						if i > end {
							break
						}
						if _, found := allTwins[i]; found {
							startInc++
							continue
						}
						allTwins[i] = true
						count += i
						startInc++
					}
				}

			}
			fmt.Printf("%v\n", allTwins)
		})

		return count
	})
}
