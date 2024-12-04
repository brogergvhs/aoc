package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening file: " + err.Error())
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic("Error reading file: " + err.Error())
	}

  xmasCount := countXMAS(grid)

	fmt.Println("The word 'XMAS' appears", xmasCount, "times.")
}

