package main

type Region struct {
	letter byte
	area   int
	sides  int
}

func CalculateTotalCost(grid [][]byte) int {
	regions := findRegions(grid)

	var totalCost int
	for _, region := range regions {
		totalCost += region.area * region.sides
	}

	return totalCost
}

func findRegions(grid [][]byte) []Region {
	rows := len(grid)
	if rows == 0 {
		return nil
	}
	cols := len(grid[0])

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var regions []Region

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if !visited[x][y] {
				letter := grid[x][y]
				region := exploreRegion2(grid, x, y, letter, visited)
				regions = append(regions, region)
			}
		}
	}

	return regions
}

func exploreRegion2(grid [][]byte, startRow, startCol int, letter byte, visited [][]bool) Region {
	queue := []Point{{startRow, startCol}}
	visited[startRow][startCol] = true

	area := 0
	sides := 0

	rows := len(grid)
	cols := len(grid[0])

	directions := []Point{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		area++

		for i, dir := range directions {
			newRow := current.x + dir.x
			newCol := current.y + dir.y

			if !isValidPoint(newRow, newCol, rows, cols) || grid[newRow][newCol] != letter {
				switch i {
				case 0: // Up
					if !isConnected(current.x-1, current.y, grid, letter, rows, cols) {
						if !isConnected(current.x, current.y+1, grid, letter, rows, cols) {
							sides++
						} else if isConnected(current.x-1, current.y+1, grid, letter, rows, cols) {
							sides++
						}
					}
				case 1: // Right
					if !isConnected(current.x, current.y+1, grid, letter, rows, cols) {
						if !isConnected(current.x+1, current.y, grid, letter, rows, cols) {
							sides++
						} else if isConnected(current.x+1, current.y+1, grid, letter, rows, cols) {
							sides++
						}
					}
				case 2: // Down
					if !isConnected(current.x+1, current.y, grid, letter, rows, cols) {
						if !isConnected(current.x, current.y-1, grid, letter, rows, cols) {
							sides++
						} else if isConnected(current.x+1, current.y-1, grid, letter, rows, cols) {
							sides++
						}
					}
				case 3: // Left
					if !isConnected(current.x, current.y-1, grid, letter, rows, cols) {
						if !isConnected(current.x-1, current.y, grid, letter, rows, cols) {
							sides++
						} else if isConnected(current.x-1, current.y-1, grid, letter, rows, cols) {
							sides++
						}
					}
				}
			} else {
				if !visited[newRow][newCol] {
					visited[newRow][newCol] = true
					queue = append(queue, Point{newRow, newCol})
				}
			}
		}
	}

	return Region{
		letter: letter,
		area:   area,
		sides:  sides,
	}
}

func isValidPoint(x, y, totalRows, totalCols int) bool {
	return x >= 0 && x < totalRows && y >= 0 && y < totalCols
}

func isConnected(x2, y2 int, grid [][]byte, letter byte, rows, cols int) bool {
	if !isValidPoint(x2, y2, rows, cols) {
		return false
	}
	return grid[x2][y2] == letter
}
