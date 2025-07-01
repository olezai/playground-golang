package main

import (
	"log"
	"os"
	"path/filepath"
	"zipdemo/zipdemo"
)

func main() {
	// Configure file paths
	filename := "nicefile.txt"
	archiveName := "nicearchive.zip"
	content := `This is some random text that will be zipped.
Hello World!
This is a demo file.`

	// Create working directory if it doesn't exist
	workDir := "output"
	if err := os.MkdirAll(workDir, 0755); err != nil {
		log.Fatalf("Failed to create working directory: %v", err)
	}

	// Use filepath.Join for cross-platform path handling
	filePath := filepath.Join(workDir, filename)
	archivePath := filepath.Join(workDir, archiveName)

	// Create and zip the file with error handling
	if err := zipdemo.CreateFile(filePath, content); err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}

	if err := zipdemo.ZipFile(archivePath, filePath); err != nil {
		log.Fatalf("Failed to create zip archive: %v", err)
	}

	log.Printf("Successfully created zip archive at: %s", archivePath)
}
