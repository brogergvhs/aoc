package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

type ClawMachine2 struct {
	ax, ay, bx, by *big.Int
	px, py         *big.Int
}

func task2(filename string) (*big.Int, int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return big.NewInt(0), 0
	}
	defer file.Close()

	var machines []ClawMachine2
	scanner := bufio.NewScanner(file)

	for {
		var ax, ay, bx, by, px, py int64
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
			machines = append(machines, ClawMachine2{
				ax: big.NewInt(ax), ay: big.NewInt(ay),
				bx: big.NewInt(bx), by: big.NewInt(by),
				px: big.NewInt(px), py: big.NewInt(py)})
		}
	}

	totalTokens := big.NewInt(0)
	prizesWon := 0

	for _, machine := range machines {
		tokens, solvable := minTokens2(machine)
		if solvable {
			totalTokens.Add(totalTokens, tokens)
			prizesWon++
		}
	}

	return totalTokens, prizesWon
}

func minTokens2(machine ClawMachine2) (*big.Int, bool) {
	ax, ay, bx, by, px, py := machine.ax, machine.ay, machine.bx, machine.by, machine.px, machine.py

	D := big.NewInt(0)
	D.Mul(ax, by)
	temp := big.NewInt(0)
	temp.Mul(ay, bx)
	D.Sub(D, temp)

	if D.Sign() == 0 {
		return nil, false
	}

	D_a := big.NewInt(0)
	D_a.Mul(px, by)
	temp.Mul(py, bx)
	D_a.Sub(D_a, temp)

	D_b := big.NewInt(0)
	D_b.Mul(ax, py)
	temp.Mul(ay, px)
	D_b.Sub(D_b, temp)

	aPresses := big.NewInt(0)
	bPresses := big.NewInt(0)
	mod := big.NewInt(0)

	aPresses.DivMod(D_a, D, mod)
	if mod.Sign() != 0 {
		return nil, false
	}

	bPresses.DivMod(D_b, D, mod)
	if mod.Sign() != 0 {
		return nil, false
	}

	if aPresses.Sign() < 0 || bPresses.Sign() < 0 {
		return nil, false
	}

	tokens := big.NewInt(0)
	tokens.Mul(aPresses, big.NewInt(3))
	tokens.Add(tokens, bPresses)

	return tokens, true
}
