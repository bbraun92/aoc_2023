package main

import (
	"aoc/util"
	"fmt"
	"slices"
	"strconv"
)

func main() {
	silverStar, goldStar := solve()
	fmt.Println("Solution to part 1: ", silverStar)
	fmt.Println("Solution to part 2: ", goldStar)
}

type GearHelper struct {
	count          int
	multiplication int
}

func solve() (int, int) {
	lines := util.ReadFileLineByLine("input.txt")

	sumParts := 0
	gears := map[string]GearHelper{}

	for y, line := range lines {
		start, end := -1, -1

		for x, characterRune := range line {
			character := string(characterRune)
			isDigit := isDigit(character)

			// Found start of number.
			if start == -1 && isDigit {
				start = x
				continue
			}

			end = x

			if !isDigit {
				end = x - 1
			}

			reachedEndOfDigits := !isDigit || x == len(line)-1

			// Found end of number.
			if start != -1 && reachedEndOfDigits {
				number, _ := strconv.Atoi(line[start : end+1])
				includeInSum := false

				// Go over -1 / +1 window of found number and check the characters.
				for i := start - 1; i <= end+1; i++ {
					if i == -1 || i == len(line) {
						continue
					}

					for j := y - 1; j <= y+1; j++ {
						if j == -1 || j == len(lines) {
							continue
						}

						candidate := string(lines[j][i])
						includeInSum = includeInSum || isSymbol(candidate)
						foundGear := isGear(candidate)

						if foundGear {
							addOrUpdateGearCandidate(gears, j, i, number)
						}
					}
				}

				if includeInSum {
					sumParts += number
				}

				start = -1
			}
		}
	}

	gearsSum := 0

	for _, value := range gears {
		if value.count == 2 {
			gearsSum += value.multiplication
		}
	}

	return sumParts, gearsSum
}

var digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func isDigit(character string) bool {
	return slices.Index(digits, character) > -1
}

func isSymbol(character string) bool {
	return character != "." && !isDigit(character)
}

func isGear(character string) bool {
	return character == "*"
}

func addOrUpdateGearCandidate(gears map[string]GearHelper, x int, y int, number int) {
	key := fmt.Sprint(x) + "-" + fmt.Sprint(y)
	gear, included := gears[key]

	if !included {
		gears[key] = GearHelper{
			count:          1,
			multiplication: number,
		}
	} else {
		gear.count++
		gear.multiplication *= number
		gears[key] = gear
	}
}
