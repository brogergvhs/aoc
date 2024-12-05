package main

func Task2(updates [][]int, rules map[int][]int) int {
	sumOfMiddlePages := 0

	for _, update := range updates {
		if !isValidOrder(update, rules) {
			reordered := reorderPages(update, rules)

			middleIndex := len(reordered) / 2
			sumOfMiddlePages += reordered[middleIndex]
		}
	}

	return sumOfMiddlePages
}

func reorderPages(update []int, rules map[int][]int) []int {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	pagesInUpdate := make(map[int]bool)

	for _, page := range update {
		pagesInUpdate[page] = true
	}

	for x, ys := range rules {
		if !pagesInUpdate[x] {
			continue
		}
		for _, y := range ys {
			if !pagesInUpdate[y] {
				continue
			}
			graph[x] = append(graph[x], y)
			inDegree[y]++
		}
	}

	// topological sort w/ Kahn's algo
	var sorted []int
	var queue []int

	for _, page := range update {
		if inDegree[page] == 0 {
			queue = append(queue, page)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		sorted = append(sorted, current)

		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sorted
}
