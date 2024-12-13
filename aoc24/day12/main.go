package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	var garden [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		garden = append(garden, []byte(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	totalCost1 := calculateTotalFencingCost(garden)
	fmt.Println("Task 1 result:", totalCost1)

  totalCost2 := CalculateTotalCost(garden)
  fmt.Println("Task 2 result:", totalCost2)
}
