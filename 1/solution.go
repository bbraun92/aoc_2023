package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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
	silverStar := solve(false)
	goldStar := solve(true)

	fmt.Println("Solution to part one: ", silverStar)
	fmt.Println("Solution to part two: ", goldStar)
}

func solve(goldStar bool) int {
	f, err := os.Open(FILE_SRC)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		firstNum, lastNum := getFirstLast(line)

		if !goldStar {
			number, _ := strconv.Atoi(firstNum.digit + lastNum.digit)
			sum += number
			continue
		}

		wordIndexFirst, wordIndexLast, wordFirst, wordLast := getWordFirstLast(line)

		first := ""
		last := ""

		if firstNum.index < wordIndexFirst {
			first = firstNum.digit
		} else {
			first = wordFirst
		}

		if lastNum.index > wordIndexLast {
			last = lastNum.digit
		} else {
			last = wordLast
		}

		number, _ := strconv.Atoi(first + last)
		sum += number
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func getFirstLast(line string) (IndexDigitPair, IndexDigitPair) {
	first := IndexDigitPair{0, ""}
	last := IndexDigitPair{0, ""}

	for i, r := range line {
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

func getWordFirstLast(line string) (int, int, string, string) {
	indexFirst := math.MaxInt
	indexLast := -1

	wordFirst := ""
	wordLast := ""

	for _, r := range words {
		index := strings.Index(line, r.word)

		if index > -1 && index < indexFirst {
			indexFirst = index
			wordFirst = r.digit
		}

		index = strings.LastIndex(line, r.word)

		if index > -1 && index > indexLast {
			indexLast = index
			wordLast = r.digit
		}
	}

	return indexFirst, indexLast, wordFirst, wordLast
}
