package main

func Task2(grid []string) int {
	simulate := func(grid []string) bool {
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

		visited := make(map[[3]int]bool)
		x, y := startX, startY

		for {
			state := [3]int{y, x, dir}
			if visited[state] {
				return true
			}
			visited[state] = true

			ny, nx := y+directions[dir][0], x+directions[dir][1]

			if ny < 0 || ny >= len(grid) || nx < 0 || nx >= len(grid[0]) {
				break
			}

			if grid[ny][nx] == '#' {
				dir = (dir + 1) % 4
			} else {
				y, x = ny, nx
			}
		}

		return false
	}

	count := 0
	for y, row := range grid {
		for x, char := range row {
			if char == '.' {
				grid[y] = row[:x] + "#" + row[x+1:]
				if simulate(grid) {
					count++
				}
				grid[y] = row
			}
		}
	}

	return count
}
