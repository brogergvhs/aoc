package main

import (
	"bufio"
	"fmt"
	"os"
)

func ParseMap(file *os.File) []string {
	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	return grid
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	grid := ParseMap(file)

	task1Result := Task1(grid)
	fmt.Println("Task 1: distinct positions visited:", task1Result)

	task2Result := Task2(grid)
	fmt.Println("Task 2: positions causing loops with extr obstacle:", task2Result)
}
