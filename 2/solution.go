package main

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	silverStar, goldStar := solve(12, 13, 14)
	fmt.Println("Solution to part one: ", silverStar)
	fmt.Println("Solution to part two: ", goldStar)
}

func solve(maxRed, maxGreen, maxBlue int) (int, int) {
	lines := util.ReadFileLineByLine("input.txt")
	sumValid, sumPow := 0, 0

	for _, line := range lines {
		withoutPrefix, _ := strings.CutPrefix(line, "Game ")
		gameString, roundsString, _ := strings.Cut(withoutPrefix, ":")

		game, _ := strconv.Atoi(gameString)
		rounds := strings.Split(roundsString, ";")

		colorToCount := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, roundString := range rounds {
			round := strings.Split(roundString, ",")

			for _, color := range round {
				colorParts := strings.Split(color, " ")
				count, _ := strconv.Atoi(colorParts[1])
				colorKey := colorParts[2]

				if count > colorToCount[colorKey] {
					colorToCount[colorKey] = count
				}
			}
		}

		if colorToCount["red"] <= maxRed && colorToCount["green"] <= maxGreen && colorToCount["blue"] <= maxBlue {
			sumValid += game
		}

		sumPow += colorToCount["red"] * colorToCount["green"] * colorToCount["blue"]
	}

	return sumValid, sumPow
}
