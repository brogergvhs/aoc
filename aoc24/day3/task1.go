package main

import (
	"regexp"
	"strconv"
)

func CalculateSum(input string) int {
	reMul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	var sum int

	matches := reMul.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])

		sum += x * y
	}

	return sum
}
