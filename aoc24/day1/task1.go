package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/brogergvhs/aoc24/utils"
)

func readInput(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var lL, rL []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		lNum, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}

		rNum, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}

		lL = append(lL, lNum)
		rL = append(rL, rNum)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return lL, rL, nil
}

func calculateTotalDistance(lL, rL []int) int {
	sort.Ints(lL)
	sort.Ints(rL)

	totalDistance := 0
	for i := 0; i < len(lL); i++ {
		diff := utils.Abs(lL[i] - rL[i])
		totalDistance += diff
	}

	return totalDistance
}
