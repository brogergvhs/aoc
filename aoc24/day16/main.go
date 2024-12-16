package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type directedPoint struct {
	point     point
	direction rune
}

const (
	north rune = '^'
	east  rune = '>'
	south rune = 'v'
	west  rune = '<'
)

type routeState struct {
	reindeer directedPoint
	steps    []directedPoint
	score    int
}

type point struct {
	x int
	y int
}

const (
	wall = '#'
	path = '.'
)

func main() {
	input, err := parseInputFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	reindeer, end, maze := parse(input)
	score, steps := findBestRoute(reindeer, end, maze)

	fmt.Printf("Task 1 result: %d\n", score)
	fmt.Printf("Task 2 result: %d\n", steps)
}

func parseInputFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func findBestRoute(reindeer directedPoint, end point, maze [][]rune) (bestScore int, totalSteps int) {
	bestScore = math.MaxInt
	queue := []routeState{{reindeer, []directedPoint{reindeer}, 0}}
	visited := make(map[directedPoint]int)
	bestRoutes := make(map[int][]directedPoint)

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if len(state.steps) > 10000 {
			continue
		}

		if state.score > bestScore {
			continue
		}

		if state.reindeer.point == end {
			if state.score <= bestScore {
				bestRouteTmp := bestRoutes[state.score]
				bestRouteTmp = append(bestRouteTmp, state.steps...)
				bestRoutes[state.score] = bestRouteTmp
				bestScore = state.score
				continue
			}
		}

		for _, candidate := range getOffsets(state.reindeer) {
			if maze[candidate.point.y][candidate.point.x] == path {
				score := state.score + 1
				if state.reindeer.direction != candidate.direction {
					score += 1000
				}

				if previous, found := visited[candidate]; found {
					if previous < score {
						continue
					}
				}
				visited[candidate] = score

				newSteps := make([]directedPoint, len(state.steps))
				copy(newSteps, state.steps)

				queue = append(queue, routeState{
					reindeer: candidate,
					steps:    append(newSteps, candidate),
					score:    score})
			}
		}
	}

	buffer := make(map[point]int)
	for _, v := range bestRoutes[bestScore] {
		buffer[v.point]++
	}

	return bestScore, len(buffer)
}

func getOffsets(reindeer directedPoint) []directedPoint {
	n := directedPoint{point{x: reindeer.point.x + 0, y: reindeer.point.y - 1}, north}
	e := directedPoint{point{x: reindeer.point.x + 1, y: reindeer.point.y + 0}, east}
	s := directedPoint{point{x: reindeer.point.x + 0, y: reindeer.point.y + 1}, south}
	w := directedPoint{point{x: reindeer.point.x - 1, y: reindeer.point.y + 0}, west}

	switch reindeer.direction {
	case north:
		return []directedPoint{n, e, w}
	case east:
		return []directedPoint{e, s, n}
	case south:
		return []directedPoint{s, w, e}
	case west:
		return []directedPoint{w, n, s}
	}

	panic("Reindeer facing unknown direction.")
}

func parse(input []string) (reindeer directedPoint, end point, maze [][]rune) {
	maze = [][]rune{}

	for y, line := range input {
		maze = append(maze, []rune{})
		for x, r := range line {
			switch r {
			case wall:
				maze[y] = append(maze[y], r)
			case path:
				maze[y] = append(maze[y], r)
			case 'S':
				reindeer = directedPoint{point{x, y}, east}
				maze[y] = append(maze[y], path)
			case 'E':
				end = point{x, y}
				maze[y] = append(maze[y], path)
			}
		}
	}

	return reindeer, end, maze
}
