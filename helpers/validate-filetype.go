package helpers

import (
	"archive/zip"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
)

// ValidateFileType checks if the file matches the expected type
func ValidateFileType(fileHeader *multipart.FileHeader, expectedType string) bool {
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return false
	}
	defer file.Close()

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return false
	}

	fileType := http.DetectContentType(buff)
	fmt.Println("Detected file type: " + fileType)

	// Map expectedType to MIME type prefix
	var mimeTypePrefix string
	switch expectedType {
	case "image":
		mimeTypePrefix = "image/"
	case "pdf":
		mimeTypePrefix = "application/pdf"
	case "epub":
		mimeTypePrefix = "application/zip"
	default:
		fmt.Printf("Unsupported file type: %s\n", expectedType)
		return false
	}

	if !strings.HasPrefix(fileType, mimeTypePrefix) {
		return false
	}

	// // Additional check for EPUB files
	// if expectedType == "epub" && mimeTypePrefix == "application/zip" {
	// 	file.Seek(0, 0) // Reset file pointer to the beginning
	// 	return isEPUB(file, fileHeader.Size)
	// }

	return true
}

// Check if the ZIP file is an EPUB
func isEPUB(file multipart.File, size int64) bool {
	zipReader, err := zip.NewReader(file, size)
	if err != nil {
		fmt.Printf("Error reading zip file: %v\n", err)
		return false
	}

	for _, f := range zipReader.File {
		if f.Name == "mimetype" {
			mimeFile, err := f.Open()
			if err != nil {
				fmt.Printf("Error opening mimetype file: %v\n", err)
				return false
			}
			defer mimeFile.Close()

			buff := make([]byte, 512)
			_, err = mimeFile.Read(buff)
			if err != nil {
				fmt.Printf("Error reading mimetype file: %v\n", err)
				return false
			}

			return strings.TrimSpace(string(buff)) == "application/epub+zip"
		}
	}

	return false
}
