package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInputFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 2 {
			leftNum, err1 := strconv.Atoi(parts[0])
			rightNum, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				return nil, nil, fmt.Errorf("invalid number in input file")
			}
			left = append(left, leftNum)
			right = append(right, rightNum)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return left, right, nil
}

func calculateSimilarityScore(left, right []int) int {
	countMap := make(map[int]int)
	for _, num := range right {
		countMap[num]++
	}

	similarityScore := 0
	for _, num := range left {
		similarityScore += num * countMap[num]
	}

	return similarityScore
}
