package main

import (
	"aoc/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type WordToDigit struct {
	word  string
	digit string
}

type IndexDigitPair struct {
	index int
	digit string
}

var words = [9]WordToDigit{
	{"one", "1"},
	{"two", "2"},
	{"three", "3"},
	{"four", "4"},
	{"five", "5"},
	{"six", "6"},
	{"seven", "7"},
	{"eight", "8"},
	{"nine", "9"},
}

const FILE_SRC = "input.txt"

func main() {
	fmt.Println("Solution to part one: ", partOne())
	fmt.Println("Solution to part two: ", partTwo())
}

func partOne() int {
	lines := util.ReadFileLineByLine(FILE_SRC)
	sum := 0

	for _, line := range lines {
		firstNum, lastNum := getNumberFirstLast(line)
		number, _ := strconv.Atoi(firstNum.digit + lastNum.digit)
		sum += number
	}

	return sum
}

func partTwo() int {
	lines := util.ReadFileLineByLine(FILE_SRC)
	sum := 0

	for _, line := range lines {
		firstNum, lastNum := getNumberFirstLast(line)
		firstWord, lastWord := getWordFirstLast(line)

		first := ""
		last := ""

		if firstNum.index < firstWord.index {
			first = firstNum.digit
		} else {
			first = firstWord.digit
		}

		if lastNum.index > lastWord.index {
			last = lastNum.digit
		} else {
			last = lastWord.digit
		}

		number, _ := strconv.Atoi(first + last)
		sum += number
	}

	return sum
}

func getNumberFirstLast(line string) (IndexDigitPair, IndexDigitPair) {
	first := IndexDigitPair{0, ""}
	last := IndexDigitPair{0, ""}

	for i, r := range line {
		// range over line returns rune. runes are unicode code points. string(r) parses back to the character behind code point as string.
		character := string(r)

		if _, err := strconv.Atoi(character); err == nil {
			first = IndexDigitPair{i, character}
			break
		}
	}

	for i := range line {
		iReversed := len(line) - 1 - i
		el := string(line[iReversed])

		if _, err := strconv.Atoi(el); err == nil {
			last = IndexDigitPair{iReversed, el}
			break
		}
	}

	return first, last
}

func getWordFirstLast(line string) (IndexDigitPair, IndexDigitPair) {
	first := IndexDigitPair{math.MaxInt, ""}
	last := IndexDigitPair{-1, ""}

	for _, r := range words {
		index := strings.Index(line, r.word)

		if index > -1 && index < first.index {
			first = IndexDigitPair{index, r.digit}
		}

		index = strings.LastIndex(line, r.word)

		if index > -1 && index > last.index {
			last = IndexDigitPair{index, r.digit}
		}
	}

	return first, last
}
