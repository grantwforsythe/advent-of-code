package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const numberOfRows = 1000

type Report struct {
	levels []int
}

func (r *Report) isSafe() bool {
	isDecreasing := r.levels[0] < r.levels[len(r.levels)-1]

	for i := 0; i < len(r.levels)-1; i++ {
		diff := r.levels[i] - r.levels[i+1]

		if isDecreasing && diff > 0 || !isDecreasing && diff < 0 {
			return false
		}

		// Convert to the absolute value
		if diff < 0 {
			diff *= -1
		}

		if diff > 3 || diff == 0 {
			return false
		}
	}

	return true
}

func main() {
	reports := parseInput()
	log.Printf("Part 1: %d", solvePart1(reports))
}

func parseInput() []Report {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Unable to open input.txt")
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	reports := make([]Report, numberOfRows)

	for i := 0; i < numberOfRows; i++ {
		row, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Failed to read line %d from input.txt", i)
		}

		levels := strings.Split(strings.Trim(row, "\n"), " ")
		report := Report{levels: make([]int, len(levels))}

		for j, level := range levels {
			value, err := strconv.Atoi(level)
			if err != nil {
				log.Fatalf("Failed to parse level at index (%d,%d)", i, j)
			}

			report.levels[j] = value
		}

		reports[i] = report
	}

	return reports
}

func solvePart1(reports []Report) int {
	sum := 0
	for _, report := range reports {
		if report.isSafe() {
			sum += 1
		}
	}
	return sum
}
