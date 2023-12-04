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
	cards := make([]int, len(lines))
	totalPoints, totalCards := 0, 0

	for i, line := range lines {
		cards[i] += 1

		game := strings.Split(line, ":")[1]
		parts := strings.Split(game, "|")

		winning := strings.Fields(parts[0])
		scratched := strings.Fields(parts[1])

		points := 0
		j := i

		for _, number := range scratched {
			if !slices.Contains(winning, number) {
				continue
			}

			if points == 0 {
				points = 1
			} else {
				points *= 2
			}

			j++
			cards[j] += cards[i]
		}

		totalPoints += points
		totalCards += cards[i]
	}

	return totalPoints, totalCards
}
