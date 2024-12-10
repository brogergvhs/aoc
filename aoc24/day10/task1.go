package main

func calculateScore(topoMap [][]int, startX, startY int) int {
	visited := make(map[Point]bool)
	queue := []Point{{startX, startY}}
	visited[Point{startX, startY}] = true
	score := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if topoMap[current.y][current.x] == 9 {
			score++
		}

		for _, dir := range directions {
			next := Point{current.x + dir.x, current.y + dir.y}
			if isValidMove(topoMap, visited, current, next) {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	return score
}

func isValidMove(topoMap [][]int, visited map[Point]bool, current, next Point) bool {
	if next.y < 0 || next.y >= len(topoMap) || next.x < 0 || next.x >= len(topoMap[0]) {
		return false
	}
	if visited[next] {
		return false
	}
	if topoMap[next.y][next.x] != topoMap[current.y][current.x]+1 {
		return false
	}
	return true
}
