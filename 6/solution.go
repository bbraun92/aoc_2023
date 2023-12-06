package main

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

const INPUT_SRC = "input.txt"

func main() {
	silverStar := partOne()
	goldStar := partTwo()

	fmt.Println("Solution to part 1: ", silverStar)
	fmt.Println("Solution to part 2: ", goldStar)
}

func partOne() int {
	times, distances := getTimesAndDistances()
	margin := 1

	for i, timeString := range times {
		time, _ := strconv.Atoi(timeString)
		distance, _ := strconv.Atoi(distances[i])
		margin *= countBeatingTimes(time, distance)
	}

	return margin
}

func partTwo() int {
	times, distances := getTimesAndDistances()

	timeString := strings.Join(times, "")
	distanceString := strings.Join(distances, "")

	time, _ := strconv.Atoi(timeString)
	distance, _ := strconv.Atoi(distanceString)

	return countBeatingTimes(time, distance)
}

func countBeatingTimes(availableTime, distanceToBeat int) int {
	beatingTimes := 0

	for holding := 1; holding < availableTime; holding++ {
		distance := (availableTime - holding) * holding

		if distance > distanceToBeat {
			beatingTimes++
		}
	}

	return beatingTimes
}

func getTimesAndDistances() ([]string, []string) {
	lines := util.ReadFileLineByLine(INPUT_SRC)
	times := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]

	return times, distances
}
