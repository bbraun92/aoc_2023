package main

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

const FILE_SRC = "input.txt"

const PREFIX = "Game "

const SUFFIX_RED = " red"
const SUFFIX_GREEN = " green"
const SUFFIX_BLUE = " blue"

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
			red, blue, green := 0, 0, 0

			for _, color := range round {

				if strings.Contains(color, SUFFIX_RED) {
					red = parseColor(color, SUFFIX_RED)
					continue
				}

				if strings.Contains(color, SUFFIX_GREEN) {
					green = parseColor(color, SUFFIX_GREEN)
					continue
				}

				if strings.Contains(color, SUFFIX_BLUE) {
					blue = parseColor(color, SUFFIX_BLUE)
					continue
				}
			}

			if red > MAX_RED || green > MAX_GREEN || blue > MAX_BLUE {
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

		red, blue, green := 0, 0, 0

		for _, roundString := range rounds {
			round := strings.Split(roundString, ",")

			for _, color := range round {

				if strings.Contains(color, SUFFIX_RED) {
					currentRed := parseColor(color, SUFFIX_RED)
					red = util.MaxInt(currentRed, red)
					continue
				}

				if strings.Contains(color, SUFFIX_GREEN) {
					currentGreen := parseColor(color, SUFFIX_GREEN)
					green = util.MaxInt(currentGreen, green)
					continue
				}

				if strings.Contains(color, SUFFIX_BLUE) {
					currentBlue := parseColor(color, SUFFIX_BLUE)
					blue = util.MaxInt(currentBlue, blue)
					continue
				}
			}
		}

		sum += (red * blue * green)
	}

	return sum
}

func parseColor(text, suffix string) int {
	trimmed, _ := strings.CutPrefix(text, " ")
	trimmed, _ = strings.CutSuffix(trimmed, suffix)
	color, _ := strconv.Atoi(trimmed)
	return color
}
