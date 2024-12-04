package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run createDay.go <dayNumber>")
		os.Exit(1)
	}

	dayNumber := os.Args[1]
	dirName := fmt.Sprintf("day%s", dayNumber)

	err := os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Printf("Failed to create directory %s: %v\n", dirName, err)
		os.Exit(1)
	}

	files := []string{"main.go", "task1.go", "task2.go", "input.txt"}

	for _, file := range files {
		filePath := filepath.Join(dirName, file)
		f, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("Failed to create file %s: %v\n", filePath, err)
			os.Exit(1)
		}
		f.Close()
	}

	fmt.Printf("Directory %s with files created successfully.\n", dirName)
}
