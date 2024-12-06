package main

func Task1(grid []string) int {
	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	dir := 0

	var startX, startY int
	for y, row := range grid {
		for x, char := range row {
			if char == '^' {
				startX, startY = x, y
				break
			}
		}
	}

	visited := make(map[[2]int]bool)
	x, y := startX, startY
	visited[[2]int{y, x}] = true

	rows, cols := len(grid), len(grid[0])

	for {
		ny, nx := y+directions[dir][0], x+directions[dir][1]

		if ny < 0 || ny >= rows || nx < 0 || nx >= cols {
			break
		}

		if grid[ny][nx] == '#' {
			dir = (dir + 1) % 4
		} else {
			y, x = ny, nx
			visited[[2]int{y, x}] = true
		}
	}

	return len(visited)
}
