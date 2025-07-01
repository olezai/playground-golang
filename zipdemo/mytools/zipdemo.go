package mytools

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func CreateFile(name string, text string) error {
	// Create a text file with random content
	content := []byte(text)
	err := os.WriteFile(name, content, 0644)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	return nil
}

func ZipFile(dfile string, sfile string) error {
	// Create a zip file
	zipFile, err := os.Create(dfile)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %v", err)
	}
	defer zipFile.Close()

	// Create a new zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Open the source file
	fileToZip, err := os.Open(sfile)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer fileToZip.Close()

	// Get file info
	info, err := fileToZip.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %v", err)
	}

	// Create zip header
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return fmt.Errorf("failed to create zip header: %v", err)
	}

	// Add file to zip
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return fmt.Errorf("failed to create zip entry: %v", err)
	}

	// Copy file contents to zip
	_, err = io.Copy(writer, fileToZip)
	if err != nil {
		return fmt.Errorf("failed to write to zip: %v", err)
	}
	return nil
}
