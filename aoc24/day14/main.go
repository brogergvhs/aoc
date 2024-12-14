package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Robot struct {
	px, py int // position
	vx, vy int // velocity
}

func parseInput(filename string) ([]Robot, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var robots []Robot
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		pParts := strings.Split(strings.TrimPrefix(parts[0], "p="), ",")
		vParts := strings.Split(strings.TrimPrefix(parts[1], "v="), ",")
		if len(pParts) != 2 || len(vParts) != 2 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		px, _ := strconv.Atoi(pParts[0])
		py, _ := strconv.Atoi(pParts[1])
		vx, _ := strconv.Atoi(vParts[0])
		vy, _ := strconv.Atoi(vParts[1])

		robots = append(robots, Robot{px, py, vx, vy})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return robots, nil
}

func main() {
	const (
		width    = 101
		height   = 103
		seconds1 = 100
		seconds2 = 1000000000
	)

	robots, err := parseInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	movedRobots := getEndRobotsPos(robots, seconds1, width, height)
	fmt.Println("Task 1 result:", calculateSafetyFactor(movedRobots, width, height))

	for i := range seconds2 {
		foundOnSec := moveRobots(robots, i, width, height)
		if foundOnSec != -1 {
			fmt.Println("Task 2 result:", foundOnSec)
			break
		}
	}
}
