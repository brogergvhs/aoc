package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	result1 := CalculateSum(input)
	fmt.Println("Result for Task 1:", result1)

	result2 := CalculateSumWithCondition(input)
	fmt.Println("Result for Task 2:", result2)
}
