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

var words = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

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

		numIndexFirst, numIndexLast, numberFirst, numberLast := getFirstLast(line)

		if !goldStar {
			number, _ := strconv.Atoi(numberFirst + numberLast)
			sum += number
			continue
		}

		wordIndexFirst, wordIndexLast, wordFirst, wordLast := getWordFirstLast(line)

		first := ""
		last := ""

		if numIndexFirst < wordIndexFirst {
			first = numberFirst
		} else {
			first = wordFirst
		}

		if numIndexLast > wordIndexLast {
			last = numberLast
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

func getFirstLast(line string) (int, int, string, string) {
	indexFirst := 0
	indexLast := 0

	first := ""
	last := ""

	for i, r := range line {
		if _, err := strconv.Atoi(string(r)); err == nil {
			indexFirst = i
			first = string(r)
			break
		}
	}

	for i := range line {
		iReversed := len(line) - 1 - i
		el := string(line[iReversed])

		if _, err := strconv.Atoi(el); err == nil {
			indexLast = iReversed
			last = el
			break
		}
	}

	return indexFirst, indexLast, first, last
}

func getWordFirstLast(line string) (int, int, string, string) {
	indexFirst := math.MaxInt
	indexLast := -1

	wordFirst := ""
	wordLast := ""

	for _, r := range words {
		index := strings.Index(line, string(r))

		if index > -1 {
			if index < indexFirst {
				indexFirst = index
				wordFirst = wordToNum(string(r))
			}
		}

		index = strings.LastIndex(line, string(r))

		if index > -1 {
			if index > indexLast {
				indexLast = index
				wordLast = wordToNum(string(r))
			}
		}
	}

	return indexFirst, indexLast, wordFirst, wordLast
}

func wordToNum(word string) string {
	switch word {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return "0"
	}
}
