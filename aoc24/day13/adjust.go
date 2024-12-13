package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func adjustInput(filename string) {
	inputFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create("updated_" + filename)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	prizePattern := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	const increment = 10000000000000

	for scanner.Scan() {
		line := scanner.Text()
		matches := prizePattern.FindStringSubmatch(line)
		if len(matches) == 3 {
			xValue, _ := strconv.ParseInt(matches[1], 10, 64)
			yValue, _ := strconv.ParseInt(matches[2], 10, 64)

			xValue += increment
			yValue += increment

			updatedLine := fmt.Sprintf("Prize: X=%d, Y=%d", xValue, yValue)
			writer.WriteString(updatedLine + "\n")
		} else {
			writer.WriteString(line + "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
	}
}
