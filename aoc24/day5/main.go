package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := make(map[int][]int)
	var updates [][]int
	isParsingUpdates := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isParsingUpdates = true
			continue
		}

		if !isParsingUpdates {
			parts := strings.Split(line, "|")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			rules[x] = append(rules[x], y)
		} else {
			parts := strings.Split(line, ",")
			update := make([]int, len(parts))
			for i, part := range parts {
				update[i], _ = strconv.Atoi(part)
			}
			updates = append(updates, update)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	task1Result := Task1(updates, rules)
	fmt.Println("Task 1 - Sum of middle pages of correctly ordered updates:", task1Result)

	task2Result := Task2(updates, rules)
	fmt.Println("Task 2 - Sum of middle pages of reordered updates:", task2Result)
}
