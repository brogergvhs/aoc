package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseOperands(input string) []int {
	parts := strings.Fields(input)
	operands := make([]int, len(parts))
	for i, p := range parts {
		operands[i], _ = strconv.Atoi(p)
	}
	return operands
}

func generateOperatorsCombinations(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}

	operators := []string{"+", "*", "||"}
	combinations := [][]string{}

	var backtrack func(int, []string)
	backtrack = func(idx int, current []string) {
		if idx == n {
			combo := make([]string, len(current))
			copy(combo, current)
			combinations = append(combinations, combo)
			return
		}
		for _, op := range operators {
			backtrack(idx+1, append(current, op))
		}
	}

	backtrack(0, []string{})
	return combinations
}

func canSolveEquation(target int, operands []int, task int) bool {
	operatorCombinations := generateOperatorsCombinations(len(operands) - 1)
	for _, combo := range operatorCombinations {
		switch task {
		case 1:
			if evaluateExpression(operands, combo) == target {
				return true
			}
		case 2:
			if evaluateExpression2(operands, combo) == target {
				return true
			}
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalCalibration := 0
	totalCalibration2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parts[0])
		operands := parseOperands(parts[1])

		if canSolveEquation(target, operands, 1) {
			totalCalibration += target
		}

		if canSolveEquation(target, operands, 2) {
			totalCalibration2 += target
		}
	}

	fmt.Println("Task 1 Result:", totalCalibration)
	fmt.Println("Task 2 Result:", totalCalibration2)
}

