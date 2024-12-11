package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func splitStone(n int) (int, int) {
	digits := strconv.Itoa(n)
	length := len(digits)
	mid := length / 2

	left, _ := strconv.Atoi(digits[:mid])
	right, _ := strconv.Atoi(digits[mid:])

	return left, right
}

func blink(stones map[int]int) map[int]int {
	newStones := make(map[int]int)

	for stone, count := range stones {
		switch {
		case stone == 0:
			newStones[1] += count
		case len(strconv.Itoa(stone))%2 == 0:
			left, right := splitStone(stone)
			newStones[left] += count
			newStones[right] += count
		default:
			newStones[stone*2024] += count
		}
	}

	return newStones
}

func simulateBlinks(initialStones []int, blinks int) int {
	stones := make(map[int]int)
	for _, stone := range initialStones {
		stones[stone]++
	}

	for i := 0; i < blinks; i++ {
		stones = blink(stones)
	}

	totalStones := 0
	for _, count := range stones {
		totalStones += count
	}

	return totalStones
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var stones []int
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		for _, num := range numbers {
			stone, _ := strconv.Atoi(num)
			stones = append(stones, stone)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	start25 := time.Now()
	totalStones25 := simulateBlinks(stones, 25)
	duration25 := time.Since(start25)

	start75 := time.Now()
	totalStones75 := simulateBlinks(stones, 75)
	duration75 := time.Since(start75)

	fmt.Printf("Task 1 result: %d (%v)\n", totalStones25, duration25)
	fmt.Printf("Task 2 result: %d (%v)\n", totalStones75, duration75)
}
