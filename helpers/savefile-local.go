package helpers

import (
	"io"
	"mime/multipart"
	"os"
)

// Save the uploaded file to the specified destination
func SaveFileLocal(fileHeader *multipart.FileHeader, dst string) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	return nil
}
