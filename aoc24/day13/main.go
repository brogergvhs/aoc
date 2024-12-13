package main

import "fmt"

func main() {
	filename := "input.txt"

	totalTokens1, prizesWon1 := task1(filename)
	fmt.Printf("Task 1 - Prizes won: %d, Total tokens: %d\n", prizesWon1, totalTokens1)

	adjustInput(filename)
	adjFilename := "updated_" + filename

	totalTokens2, prizesWon2 := task2(adjFilename)
	fmt.Printf("Task 2 - Prizes won: %d, Total tokens: %s\n", prizesWon2, totalTokens2.String())
}
