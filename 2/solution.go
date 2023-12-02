package main

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

const FILE_SRC = "input.txt"
const PREFIX = "Game "

func main() {
	fmt.Println("Solution to part one: ", partOne())
	fmt.Println("Solution to part two: ", partTwo())
}

func partOne() int {
	const MAX_RED = 12
	const MAX_GREEN = 13
	const MAX_BLUE = 14

	lines := util.ReadFileLineByLine(FILE_SRC)
	sumValid := 0

	for _, line := range lines {
		withoutPrefix, _ := strings.CutPrefix(line, PREFIX)
		gameString, roundsString, _ := strings.Cut(withoutPrefix, ":")

		game, _ := strconv.Atoi(gameString)
		rounds := strings.Split(roundsString, ";")

		gameValid := true

		for _, roundString := range rounds {
			round := strings.Split(roundString, ",")
			colorToCount := make(map[string]int)

			for _, color := range round {
				colorParts := strings.Split(color, " ")
				count, _ := strconv.Atoi(colorParts[1])
				colorToCount[colorParts[2]] = count
			}

			if colorToCount["red"] > MAX_RED || colorToCount["green"] > MAX_GREEN || colorToCount["blue"] > MAX_BLUE {
				gameValid = false
				break
			}
		}

		if gameValid {
			sumValid += game
		}
	}

	return sumValid
}

func partTwo() int {
	lines := util.ReadFileLineByLine(FILE_SRC)
	sum := 0

	for _, line := range lines {
		withoutPrefix, _ := strings.CutPrefix(line, PREFIX)
		_, roundsString, _ := strings.Cut(withoutPrefix, ":")
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

		sum += colorToCount["red"] * colorToCount["green"] * colorToCount["blue"]
	}

	return sum
}
