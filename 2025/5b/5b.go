package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/francois76/adventOfCode/shared"
)

func main() {
	shared.Run(func() any {
		count := 0
		allRanges := []ranges{}
		finished := false
		shared.Open("../5/5.txt", func(fileScanner *bufio.Scanner) {
			line := fileScanner.Text()
			if line == "" || finished {
				finished = true
				return
			}

			allRanges = append(allRanges, buildRange(line))
		})

		// Supprime les intervals inclus dans d'autres
		removeIncluded := func(r []ranges) []ranges {
			toKeep := make([]bool, len(r))
			for i := range toKeep {
				toKeep[i] = true
			}

			for i := range r {
				if !toKeep[i] {
					continue
				}
				for j := range r {
					if i == j || !toKeep[j] {
						continue
					}
					// Si r[j] est inclus dans r[i], on marque r[j] pour suppression
					if r[j].min >= r[i].min && r[j].max <= r[i].max {
						fmt.Println("removing ", r[j], " (included in ", r[i], ")")
						toKeep[j] = false
					}
				}
			}

			newRanges := []ranges{}
			for i := range r {
				if toKeep[i] {
					newRanges = append(newRanges, r[i])
				}
			}
			return newRanges
		}

		// Merge les intervals qui se chevauchent
		mergeOverlapping := func(r []ranges) []ranges {
			newRanges := []ranges{}
			merged := make([]bool, len(r))

			for i := range r {
				if merged[i] {
					continue
				}
				current := r[i]
				for j := i + 1; j < len(r); j++ {
					if merged[j] {
						continue
					}
					if newRange, ok := current.merge(r[j]); ok {
						fmt.Println("merging ", current, " and ", r[j], " into ", newRange)
						current = newRange
						merged[j] = true
					}
				}
				newRanges = append(newRanges, current)
			}
			return newRanges
		}

		// Boucle jusqu'à stabilisation
		for {
			allRanges = removeIncluded(allRanges)
			newRanges := mergeOverlapping(allRanges)
			if fmt.Sprintf("%v", newRanges) == fmt.Sprintf("%v", allRanges) {
				break
			}
			allRanges = newRanges
		}

		for _, r := range allRanges {
			count += int(r.max - r.min + 1)
		}
		fmt.Println(allRanges)

		return count
	})
}

type ranges struct {
	min int64
	max int64
}

func (r ranges) merge(other ranges) (ranges, bool) {
	// Pas de chevauchement
	if r.max < other.min || other.max < r.min {
		return ranges{}, false
	}

	// Si l'un est complètement inclus dans l'autre, pas de merge
	// (on les gère séparément)
	if (r.min >= other.min && r.max <= other.max) || (other.min >= r.min && other.max <= r.max) {
		return ranges{}, false
	}

	// Merge avec chevauchement partiel
	result := ranges{}
	if r.min < other.min {
		result.min = r.min
	} else {
		result.min = other.min
	}
	if r.max > other.max {
		result.max = r.max
	} else {
		result.max = other.max
	}
	return result, true
}

func buildRange(line string) ranges {
	items := strings.Split(line, "-")
	min, _ := strconv.ParseInt(items[0], 10, 64)
	max, _ := strconv.ParseInt(items[1], 10, 64)

	return ranges{min: min, max: max}
}
