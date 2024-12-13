package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ClawMachine1 struct {
	ax, ay, bx, by, px, py int
}

func task1(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, 0
	}
	defer file.Close()

	var machines []ClawMachine1
	scanner := bufio.NewScanner(file)

	for {
		var ax, ay, bx, by, px, py int
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if strings.HasPrefix(line, "Button A:") {
			fmt.Sscanf(line, "Button A: X+%d, Y+%d", &ax, &ay)
			if !scanner.Scan() {
				break
			}
			line = scanner.Text()
			fmt.Sscanf(line, "Button B: X+%d, Y+%d", &bx, &by)
			if !scanner.Scan() {
				break
			}
			line = scanner.Text()
			fmt.Sscanf(line, "Prize: X=%d, Y=%d", &px, &py)
			machines = append(machines, ClawMachine1{ax, ay, bx, by, px, py})
		}
	}

	totalTokens := 0
	prizesWon := 0

	for _, machine := range machines {
		tokens, solvable := minTokens1(machine)
		if solvable {
			totalTokens += tokens
			prizesWon++
		}
	}

	return totalTokens, prizesWon
}

func minTokens1(machine ClawMachine1) (int, bool) {
	ax, ay, bx, by, px, py := machine.ax, machine.ay, machine.bx, machine.by, machine.px, machine.py

	minTokens := int(^uint(0) >> 1) // Max int value
	foundSolution := false

	for aPresses := 0; aPresses <= 100; aPresses++ {
		for bPresses := 0; bPresses <= 100; bPresses++ {
			if aPresses*ax+bPresses*bx == px && aPresses*ay+bPresses*by == py {
				tokens := aPresses*3 + bPresses
				if tokens < minTokens {
					minTokens = tokens
					foundSolution = true
				}
			}
		}
	}

	return minTokens, foundSolution
}
