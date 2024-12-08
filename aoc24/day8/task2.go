package main

func Task2(board Board[cell], antennas map[rune][]Position) int {
	antinodes := make(map[Position]struct{})
	for _, positions := range antennas {
		if len(positions) <= 1 {
			continue
		}
		findAntinodes2(board, positions, antinodes)
		for _, pos := range positions {
			antinodes[pos] = struct{}{}
		}
	}

	return len(antinodes)
}

func findAntinodes2(board Board[cell], positions []Position, antinodes map[Position]struct{}) {
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			x := positions[i]
			y := positions[j]
			setAntinode2(board, x, y, antinodes)
			setAntinode2(board, y, x, antinodes)
		}
	}
}

func setAntinode2(board Board[cell], x, y Position, antinodes map[Position]struct{}) {
	d := posDiff(x, y)
	z := posDiff(y, d)
	if board.Contains(z) {
		antinodes[z] = struct{}{}
		setAntinode2(board, y, z, antinodes)
	}
}
