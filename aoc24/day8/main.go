package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	Row int
	Col int
}

type cell struct {
	emt  bool
	freq rune
}

type Board[T any] struct {
	Positions map[Position]T
	MaxRows   int
	MinRows   int
	MaxCols   int
	MinCols   int
}

func (b Board[T]) Contains(position Position) bool {
	return position.Row < b.MaxRows &&
		position.Row >= b.MinRows &&
		position.Col < b.MaxCols &&
		position.Col >= b.MinCols
}

func ParseBoard[T any](lines []string, fn func(r rune, pos Position) T) Board[T] {
	positions := make(map[Position]T, len(lines)*len(lines[0]))
	for row, line := range lines {
		for col, c := range line {
			pos := Position{Row: row, Col: col}
			positions[pos] = fn(c, pos)
		}
	}
	return Board[T]{
		Positions: positions,
		MaxRows:   len(lines),
		MinRows:   0,
		MaxCols:   len(lines[0]),
		MinCols:   0,
	}
}

func posDiff(x, y Position) Position {
	return Position{
		Row: x.Row - y.Row,
		Col: x.Col - y.Col,
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	antennas := make(map[rune][]Position)
	board := ParseBoard(res, func(r rune, pos Position) cell {
		if r == '.' {
			return cell{emt: true}
		}
		antennas[r] = append(antennas[r], pos)
		return cell{freq: r}
	})

	fmt.Println("Task 1 result:", Task1(board, antennas))
	fmt.Println("Task 2 result:", Task2(board, antennas))
}
