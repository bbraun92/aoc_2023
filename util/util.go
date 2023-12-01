package util

import (
	"bufio"
	"log"
	"os"
)

func ReadFileLineByLine(fileSrc string) []string {
	f, err := os.Open(fileSrc)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
