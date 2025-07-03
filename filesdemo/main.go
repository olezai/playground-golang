package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Create a new directory
	err := os.Mkdir("demo-dir", 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	// Create and write to a file
	content := []byte("Hello, this is a demo file!\nWelcome to file handling in Go.")
	err = os.WriteFile("demo-dir/test.txt", content, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	// Read file contents
	data, err := os.ReadFile("demo-dir/test.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("File contents:\n%s\n", string(data))

	// List directory contents
	files, err := os.ReadDir("demo-dir")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	fmt.Println("\nDirectory contents:")
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			fmt.Println("Error getting file info:", err)
			continue
		}
		fmt.Printf("Name: %s, Size: %d bytes, IsDir: %v\n",
			file.Name(), info.Size(), file.IsDir())
	}

	// Walk through directory
	fmt.Println("\nWalking through directory tree:")
	err = filepath.Walk("demo-dir", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("Path: %s\n", path)
		return nil
	})

	// Append to file
	f, err := os.OpenFile("demo-dir/test.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	newContent := []byte("\nThis line was appended!")
	if _, err := f.Write(newContent); err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}

	// Clean up - remove directory and its contents
	err = os.RemoveAll("demo-dir")
	if err != nil {
		fmt.Println("Error removing directory:", err)
		return
	}
}
