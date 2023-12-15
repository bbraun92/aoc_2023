package main

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	silverStar := partOne()
	goldStar := partTwo()

	fmt.Println("Solution to part one:", silverStar)
	fmt.Println("Solution to part two:", goldStar)
}

func partOne() int {
	lines := util.ReadFileLineByLine("input.txt")
	total := 0

	for _, line := range lines {
		history := getHistory(line)
		sequences := getSequences(history)

		sum := 0

		for _, sequence := range sequences {
			sum += sequence[len(sequence)-1]
		}

		total += sum
	}

	return total
}

func partTwo() int {
	lines := util.ReadFileLineByLine("input.txt")
	total := 0

	for _, line := range lines {
		history := getHistory(line)
		sequences := getSequences(history)

		sum := 0

		for i := len(sequences) - 1; i > 0; i-- {
			sum = -sum + sequences[i-1][0]
		}

		total += sum
	}

	return total
}

func getHistory(line string) []int {
	historyStrings := strings.Fields(line)
	history := make([]int, len(historyStrings))

	for i, number := range historyStrings {
		history[i], _ = strconv.Atoi(number)
	}

	return history
}

func getSequences(history []int) [][]int {
	sequences := make([][]int, 0)
	sequences = append(sequences, history)
	previousSequence := history

	for true {
		sequence := make([]int, len(previousSequence)-1)

		onlyZeros := true

		for i := 0; i < len(sequence); i++ {
			sequence[i] = previousSequence[i+1] - previousSequence[i]

			if sequence[i] != 0 {
				onlyZeros = false
			}
		}

		sequences = append(sequences, sequence)

		if onlyZeros {
			break
		}

		previousSequence = sequence
	}

	return sequences
}
