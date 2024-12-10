package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x, y int
}

var directions = []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	var topoMap [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, ch := range line {
			row[i], _ = strconv.Atoi(string(ch))
		}
		topoMap = append(topoMap, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	res1 := 0
	res2 := 0
	for y := 0; y < len(topoMap); y++ {
		for x := 0; x < len(topoMap[0]); x++ {
			if topoMap[y][x] == 0 {
				res1 += calculateScore(topoMap, x, y)
				res2 += calculateRating(topoMap, x, y)
			}
		}
	}

	fmt.Println("Task 1 result:", res1)
	fmt.Println("Task 2 result:", res2)
}
