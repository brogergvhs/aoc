package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type InfoPart struct {
	id      int
	empty   bool
	isMoved bool
}

func parseInput(input string) []InfoPart {
	var disk []InfoPart
	fileID := 0
	file := true

  for _, ch := range input {
		length, _ := strconv.Atoi(string(ch))
		if file {
			for i := 0; i < length; i++ {
				disk = append(disk, InfoPart{id: fileID})
			}
			fileID++
		} else {
			for i := 0; i < length; i++ {
				disk = append(disk, InfoPart{empty: true})
			}
		}
		file = !file
	}
	return disk
}

func calculateChecksum(disk []InfoPart) int {
	checksum := 0
	for i, c := range disk {
		if !c.empty {
			checksum += i * c.id
		}
	}
	return checksum
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	disk := parseInput(input)

	compressedDisk1 := compressLayout1(disk)
	checksum1 := calculateChecksum(compressedDisk1)
	fmt.Println("Task 1 result:", checksum1)

	compressedDisk2 := compressLayout2(disk)
	checksum2 := calculateChecksum(compressedDisk2)
	fmt.Println("Task 2 result:", checksum2)
}
