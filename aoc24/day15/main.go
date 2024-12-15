package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	r := strings.NewReplacer("#", "##", "O", "[]", ".", "..", "@", "@.")

	task1 := run(split[0], split[1])
	task2 := run(r.Replace(split[0]), split[1])

	fmt.Println("Task 1 result:", task1)
	fmt.Println("Task 2 result:", task2)
}

func run(warehouse, moves string) int {
	robot := Point{}
	grid := map[Point]rune{}
	for y, slice := range strings.Fields(warehouse) {
		for x, char := range slice {
			if char == '@' {
				robot = Point{x, y}
				char = '.'
			}
			grid[Point{x, y}] = char
		}
	}

	dirs := map[rune]Point{
		'^': {0, -1}, '>': {1, 0}, 'v': {0, 1}, '<': {-1, 0},
		'[': {1, 0}, ']': {-1, 0},
	}

robotHandler:
	for _, char := range strings.ReplaceAll(moves, "\n", "") {
		queue, boxes := []Point{robot}, map[Point]rune{}
		for len(queue) > 0 {
			point := queue[0]
			queue = queue[1:]

			if _, ok := boxes[point]; ok {
				continue
			}
			boxes[point] = grid[point]

			switch n := point.Add(dirs[char]); grid[n] {
			case '#':
				continue robotHandler
			case '[', ']':
				queue = append(queue, n.Add(dirs[grid[n]]))
				fallthrough
			case 'O':
				queue = append(queue, n)
			}
		}

		for box := range boxes {
			grid[box] = '.'
		}
		for box := range boxes {
			grid[box.Add(dirs[char])] = boxes[box]
		}
		robot = robot.Add(dirs[char])
	}

	gps := 0
	for point, entity := range grid {
		if entity == 'O' || entity == '[' {
			gps += 100*point.y + point.x
		}
	}
	return gps
}
