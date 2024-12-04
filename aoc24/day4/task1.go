package main

const word1 = "XMAS"

var dirs1 = [][2]int{
	{0, 1},  // horiz (right)
	{1, 0},  // vert (down)
	{1, 1},  // diag (down-right)
	{1, -1}, // diag (down-left)
	{0, -1}, // horiz (left)
	{-1, 0}, // vert (up)
	{-1, -1}, // diag (up-left)
	{-1, 1}, // diag (up-right)
}

func countXMAS(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])

	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range dirs1 {
				if checkWord1(grid, r, c, dir[0], dir[1], word1) {
					count++
				}
			}
		}
	}

	return count
}

func checkWord1(grid []string, r, c, dr, dc int, word string) bool {
	for i := 0; i < len(word); i++ {
		newR := r + dr*i
		newC := c + dc*i

		if newR < 0 || newR >= len(grid) || newC < 0 || newC >= len(grid[0]) {
			return false
		}

		if grid[newR][newC] != word[i] {
			return false
		}
	}

	return true
}

