package main

func Task1(board Board[cell], antennas map[rune][]Position) int {
	antinodes := make(map[Position]struct{})
	for _, positions := range antennas {
		findAntinodes(board, positions, antinodes)
	}

	return len(antinodes)
}

func findAntinodes(board Board[cell], positions []Position, antinodes map[Position]struct{}) {
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			x := positions[i]
			y := positions[j]
			setAntinode(board, x, y, antinodes)
			setAntinode(board, y, x, antinodes)
		}
	}
}

func setAntinode(board Board[cell], x, y Position, antinodes map[Position]struct{}) {
	d := posDiff(x, y)
	z := posDiff(y, d)
	if board.Contains(z) {
		antinodes[z] = struct{}{}
	}
}
