package main

import "fmt"

func moveRobots(robots []Robot, seconds, width, height int) int {
	movedRobots := make([]Robot, len(robots))
	copy(movedRobots, robots)

	for i := range movedRobots {
		r := &movedRobots[i]
		r.px = (r.px + r.vx*seconds) % width
		if r.px < 0 {
			r.px += width
		}

		r.py = (r.py + r.vy*seconds) % height
		if r.py < 0 {
			r.py += height
		}
	}

	return generateMap(movedRobots, width, height, seconds)
}

func generateMap(robots []Robot, width, height, curSec int) int {
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for _, robot := range robots {
		if robot.px >= 0 && robot.px < width && robot.py >= 0 && robot.py < height {
			grid[robot.py][robot.px] = '#'
		}
	}

	if checkTrianglesWithBase(grid) {
		return curSec
	}

	return -1
}

func printMap(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func checkTrianglesWithBase(grid [][]rune) bool {
	width := len(grid[0])
	height := len(grid)

	for y := 0; y < height; y++ {
		baseStart := -1
		baseLength := 0

		for x := 0; x < width; x++ {
			if grid[y][x] == '#' {
				if baseStart == -1 {
					baseStart = x
				}
				baseLength++
			} else {
				baseStart = -1
				baseLength = 0
			}

			if baseLength >= 7 {
				if isTriangle(grid, y, baseStart, baseLength, -1) {
					printMap(grid)
					return true
				}
			}
		}
	}

	return false
}

func isTriangle(grid [][]rune, baseY, baseStart, baseLength, direction int) bool {
	width := len(grid[0])
	height := len(grid)

	for level := 1; level <= baseLength/2; level++ {
		startX := baseStart + level
		endX := baseStart + baseLength - level - 1
		currentY := baseY + level*direction

		if currentY < 0 || currentY >= height {
			return false
		}

		for x := startX; x <= endX; x++ {
			if x < 0 || x >= width || grid[currentY][x] != '#' {
				return false
			}
		}
	}

	return true
}
