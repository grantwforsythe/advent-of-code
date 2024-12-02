package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const numberOfRows = 1000

func main() {
	leftSide, rightSide := parseInput()
	log.Printf("Part 1: %d", solvePart1(leftSide, rightSide))
	log.Printf("Part 2: %d", solvePart2(leftSide, rightSide))
}

func parseInput() ([]int, []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Unable to open input.txt")
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	leftSide := make([]int, numberOfRows)
	rightSide := make([]int, numberOfRows)

	for i := 0; i < numberOfRows; i++ {
		row, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Failed to read line %d from input.txt", i)
		}

		values := strings.Split(strings.Trim(row, "\n"), "   ")

		leftSide[i], err = strconv.Atoi(values[0])
		if err != nil {
			log.Fatalf("Failed to parse leftSide at index %d", i)
		}

		rightSide[i], err = strconv.Atoi(values[1])
		if err != nil {
			log.Fatalf("Failed to parse rigthSide at index %d", i)
		}
	}

	// TODO: Sort in decending order
	slices.Sort(leftSide)
	slices.Sort(rightSide)

	return leftSide, rightSide
}

func solvePart1(leftSide, rigthSide []int) int {
	// Calculate difference
	sum := 0
	for i := 0; i < numberOfRows; i++ {
		diff := leftSide[numberOfRows-1-i] - rigthSide[numberOfRows-1-i]

		if diff < 0 {
			sum += diff * -1
		} else {
			sum += diff
		}
	}

	return sum
}

func solvePart2(leftSide, rightSide []int) int {
	frequency := make(map[int]int)
	for _, value := range rightSide {
		frequency[value] += 1
	}

	sum := 0
	for _, value := range leftSide {
		sum += value * frequency[value]
	}

	return sum
}
