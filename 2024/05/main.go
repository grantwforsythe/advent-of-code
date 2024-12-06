// Day 5 is a sorting problem
package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type PageOrderRule struct {
	X int // Has to come before Y
	Y int // Has to come after X
}

func main() {
	rules, pageOrderings := parseInput("input.txt")
	log.Print(rules[0])
	log.Print(pageOrderings[0])
}

func parseInput(filename string) ([]*PageOrderRule, [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open %s", filename)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	// PERF: Not memory efficient
	rules := make([]*PageOrderRule, 0)
	pageOrderings := make([][]int, 0)

	isRulesSection := true

	for {
		row, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Failed to read a row")
		}

		// No longer reading rules
		if row == "\n" {
			isRulesSection = false
			continue
		}

		if isRulesSection {
			rule, err := parseRule(row)
			if err != nil {
				log.Fatalf("Failed to parse %s", row)
			}
			rules = append(rules, rule)
		} else {
			pageOrdering, err := parsePageOrdering(row)
			if err != nil {
				log.Fatalf("Failed to parse %s", row)
			}
			pageOrderings = append(pageOrderings, pageOrdering)
		}
	}

	return rules, pageOrderings
}

func parseRule(row string) (*PageOrderRule, error) {
	pageNumbers := strings.Split(strings.TrimSpace(row), "|")

	x, err := strconv.Atoi(pageNumbers[0])
	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(pageNumbers[1])
	if err != nil {
		return nil, err
	}

	return &PageOrderRule{x, y}, nil
}

func parsePageOrdering(row string) ([]int, error) {
	pageNumbers := strings.Split(strings.TrimSpace(row), ",")

	pageOrdering := make([]int, 0)
	for _, pageNumber := range pageNumbers {
		value, err := strconv.Atoi(pageNumber)
		if err != nil {
			return nil, err
		}

		pageOrdering = append(pageOrdering, value)
	}

	return pageOrdering, nil
}
