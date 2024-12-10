package main

func calculateRating(topoMap [][]int, startX, startY int) int {
	visited := make(map[Point]bool)
	return dfs(topoMap, visited, startX, startY)
}

func dfs(topoMap [][]int, visited map[Point]bool, x, y int) int {
	if x < 0 || y < 0 || y >= len(topoMap) || x >= len(topoMap[0]) {
		return 0
	}
	if visited[Point{x, y}] {
		return 0
	}

	visited[Point{x, y}] = true
	rating := 0
	if topoMap[y][x] == 9 {
		rating = 1
	}

	for _, dir := range directions {
		nextX, nextY := x+dir.x, y+dir.y
		if isValidTrail2(topoMap, x, y, nextX, nextY) {
			rating += dfs(topoMap, visited, nextX, nextY)
		}
	}

	visited[Point{x, y}] = false
	return rating
}

func isValidTrail2(topoMap [][]int, x, y, nextX, nextY int) bool {
	if nextX < 0 || nextY < 0 || nextY >= len(topoMap) || nextX >= len(topoMap[0]) {
		return false
	}
	if topoMap[nextY][nextX] != topoMap[y][x]+1 {
		return false
	}
	return true
}
