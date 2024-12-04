package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Prod struct {
	X int
	Y int
}

func (p *Prod) mul() int {
	return p.X * p.Y
}

func main() {
	prods := parseInput()
	log.Printf("Part 1: %d", solvePart1(prods))
}

func parseInput() []Prod {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Unable to open input.txt")
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	r, err := regexp.Compile(`mul\((\d+),(\d+)\)`)
	if err != nil {
		log.Fatal("Unable to compile the regex")
	}

	// PERF: Memory inefficient as the slice constantly needs to reallocate memory
	prods := make([]Prod, 0)

	for {
		row, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Failed to read line from input.txt")
		}

		matches := r.FindAllStringSubmatch(strings.Trim(row, "\n"), -1)
		for _, match := range matches {
			x, err := strconv.Atoi(strings.TrimSpace(match[1]))
			if err != nil {
				log.Fatalf("Failed to parse %s", match[1])
			}

			y, err := strconv.Atoi(strings.TrimSpace(match[2]))
			if err != nil {
				log.Fatalf("Failed to parse %s", match[2])
			}

			prods = append(prods, Prod{X: x, Y: y})
		}
	}

	return prods
}

func solvePart1(prods []Prod) int {
	product := 0
	for _, prod := range prods {
		product += prod.X * prod.Y
	}

	return product
}
