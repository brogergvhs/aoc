package main

func calculateTotalFencingCost(garden [][]byte) int {
	height := len(garden)
	width := len(garden[0])

	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}

	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var totalCost int

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if !visited[i][j] {
				typeOfPlant := garden[i][j]
				area, sides := exploreRegion(garden, visited, Point{i, j}, typeOfPlant, directions)
				totalCost += area * sides
			}
		}
	}

	return totalCost
}

func exploreRegion(garden [][]byte, visited [][]bool, start Point, typeOfPlant byte, directions []Point) (int, int) {
	queue := []Point{start}
	visited[start.x][start.y] = true

	height := len(garden)
	width := len(garden[0])
	area := 0
	sides := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		area++

		for _, d := range directions {
			nx, ny := current.x+d.x, current.y+d.y

			if nx < 0 || nx >= height || ny < 0 || ny >= width || garden[nx][ny] != typeOfPlant {
				sides++
			} else if !visited[nx][ny] {
				visited[nx][ny] = true
				queue = append(queue, Point{nx, ny})
			}
		}
	}

	return area, sides
}
