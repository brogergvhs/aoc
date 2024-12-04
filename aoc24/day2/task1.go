package main

import (
	"github.com/brogergvhs/aoc24/utils"
)

func isSafe(report []int) bool {
	increasing := report[1] > report[0]
	decreasing := report[1] < report[0]

	for i := 0; i < len(report)-1; i++ {
		diff := utils.Abs(report[i+1] - report[i])
		if diff < 1 || diff > 3 {
			return false
		}
		if (increasing && report[i+1] <= report[i]) || (decreasing && report[i+1] >= report[i]) {
			return false
		}
	}

	return true
}
