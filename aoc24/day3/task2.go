package main

import (
	"regexp"
	"strconv"
)

func CalculateSumWithCondition(input string) int {
	reMul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	reDoDont := regexp.MustCompile(`do\(\)|don't\(\)`)

	var sum int

	enabled := true

	matches := reDoDont.FindAllStringIndex(input, -1)

	lastEnd := 0
	for _, match := range matches {
		chunk := input[lastEnd:match[0]]

		mulMatches := reMul.FindAllStringSubmatch(chunk, -1)
		for _, mul := range mulMatches {
			if enabled {
				x, _ := strconv.Atoi(mul[1])
				y, _ := strconv.Atoi(mul[2])
				sum += x * y
			}
		}

		instruction := input[match[0]:match[1]]
		if instruction == "do()" {
			enabled = true
		} else if instruction == "don't()" {
			enabled = false
		}

		lastEnd = match[1]
	}

	chunk := input[lastEnd:]
	mulMatches := reMul.FindAllStringSubmatch(chunk, -1)
	for _, mul := range mulMatches {
		if enabled {
			x, _ := strconv.Atoi(mul[1])
			y, _ := strconv.Atoi(mul[2])
			sum += x * y
		}
	}

	return sum
}

