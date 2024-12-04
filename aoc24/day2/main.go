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
	safeReportCount := 0
	safeReportCountV2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		report := make([]int, len(parts))

		for i, part := range parts {
			report[i], err = strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error parsing number:", err)
				return
			}
		}

		if isSafe(report) {
			safeReportCount++
		}

		if canBecomeSafeByRemovingOne(report) {
			safeReportCountV2++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}

	fmt.Printf("Number of safe reports: %d\n", safeReportCount)
	fmt.Printf("Number of safe reports (v2): %d\n", safeReportCountV2)
}
