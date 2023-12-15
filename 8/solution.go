package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

type Directions struct {
	left  string
	right string
}

func main() {
	silverStar := partOne()
	goldStar := partTwo()

	fmt.Println("Solution to part one: ", silverStar)
	fmt.Println("Solution to part two: ", goldStar)
}

func partOne() int {
	lines := util.ReadFileLineByLine("input.txt")

	instructions := strings.Split(lines[0], "")
	placeToDirections := getLocationToDirectionsMap(lines[2:])

	current := "AAA"
	at := 0
	steps := 0

	for current != "ZZZ" {
		steps++

		if instructions[at] == "L" {
			current = placeToDirections[current].left
		}

		if instructions[at] == "R" {
			current = placeToDirections[current].right
		}

		at++

		if at == len(instructions) {
			at = 0
		}
	}

	return steps
}

func partTwo() int {
	lines := util.ReadFileLineByLine("input.txt")

	instructions := strings.Split(lines[0], "")
	placeToDirections := getLocationToDirectionsMap(lines[2:])

	currents := make([]string, 0)

	for i := range placeToDirections {
		if string(i[2]) == "A" {
			currents = append(currents, i)
		}
	}

	searching := true
	at := 0
	steps := 0
	found := make([]int, len(currents))

	for searching {
		steps++

		for i, current := range currents {
			if found[i] > 0 {
				continue
			}

			if instructions[at] == "L" {
				currents[i] = placeToDirections[current].left
			}

			if instructions[at] == "R" {
				currents[i] = placeToDirections[current].right
			}

			if string(currents[i][2]) == "Z" {
				found[i] = steps
			}
		}

		allFound := true

		for _, x := range found {
			if x == 0 {
				allFound = false
			}
		}

		if allFound {
			searching = false
			break
		}

		at++

		if at == len(instructions) {
			at = 0
		}
	}

	solution := leastCommonMultiple(found[0], found[1])

	for _, next := range found[2:] {
		solution = leastCommonMultiple(solution, next)
	}

	return solution
}

func getLocationToDirectionsMap(lines []string) map[string]Directions {
	placeToDirections := map[string]Directions{}

	for _, line := range lines {
		parts := strings.Fields(line)

		from := parts[0]
		left := parts[2][1:4]
		right := parts[3][:3]

		placeToDirections[from] = Directions{
			left:  left,
			right: right,
		}
	}

	return placeToDirections
}

// Euclid's algorithm for gcd
func greatestCommonDivisor(x, y int) int {
	if y == 0 {
		return x
	}

	if x == 0 {
		return y
	}

	if x > y {
		return greatestCommonDivisor(x-y, y)
	} else {
		return greatestCommonDivisor(x, y-x)
	}
}

// gcd(x, y) * lcm(x, y) = x * y for x, y >= 0, x, y being integers
// => lcm(x, y) = (x * y) / gcd(x, y)
func leastCommonMultiple(x, y int) int {
	return (x * y) / greatestCommonDivisor(x, y)
}
