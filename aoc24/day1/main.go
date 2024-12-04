package main

import (
	"fmt"
)

func main() {
	inputFile := "input.txt"

	lL1, rL1, err := readInput(inputFile)
	if err != nil {
		fmt.Printf("Error reading input for Task 1: %v\n", err)
		return
	}

	if len(lL1) != len(rL1) {
		fmt.Println("Error: The two lists must have the same number of elements.")
		return
	}

	totalDistance := calculateTotalDistance(lL1, rL1)
	fmt.Printf("Task 1 - Total Distance: %d\n", totalDistance)

	lL2, rL2, err := parseInputFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input for Task 2: %v\n", err)
		return
	}
	similarityScore := calculateSimilarityScore(lL2, rL2)
	fmt.Printf("Task 2 - Similarity Score: %d\n", similarityScore)
}
