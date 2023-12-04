package main

import (
	"aoc/util"
	"fmt"
	"slices"
	"strings"
)

func main() {
	silverStar, goldStar := solve()
	fmt.Println("Solution to part 1: ", silverStar)
	fmt.Println("Solution to part 2: ", goldStar)
}

func solve() (int, int) {
	lines := util.ReadFileLineByLine("input.txt")
	points := 0
	totalCards := 0
	bonusCards := make([]int, len(lines))

	for i, line := range lines {
		bonusCards[i] += 1

		cards := strings.Split(line, ":")[1]
		parts := strings.Split(cards, "|")

		winning := strings.Fields(parts[0])
		scratched := strings.Fields(parts[1])

		val := 0
		j := i

		for _, number := range scratched {
			if !slices.Contains(winning, number) {
				continue
			}

			if val == 0 {
				val = 1
			} else {
				val *= 2
			}

			j++
			bonusCards[j] += bonusCards[i]
		}

		totalCards += bonusCards[i]
		points += val
	}

	return points, totalCards
}
