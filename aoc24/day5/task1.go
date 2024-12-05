package main

func Task1(updates [][]int, rules map[int][]int) int {
	sumOfMiddlePages := 0

	for _, update := range updates {
		if isValidOrder(update, rules) {
			middleIndex := len(update) / 2
			sumOfMiddlePages += update[middleIndex]
		}
	}

	return sumOfMiddlePages
}

func isValidOrder(update []int, rules map[int][]int) bool {
	position := make(map[int]int)

	for i, page := range update {
		position[page] = i
	}

	for x, ys := range rules {
		for _, y := range ys {
			posX, hasX := position[x]
			posY, hasY := position[y]
			if hasX && hasY && posX > posY {
				return false
			}
		}
	}

	return true
}
