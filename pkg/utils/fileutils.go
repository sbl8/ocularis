package utils

import (
	"os"
)

// ReadFile reads a file and returns its content as a string.
func ReadFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// WriteFile writes content to a file.
func WriteFile(filePath, content string) error {
	return os.WriteFile(filePath, []byte(content), 0644)
}
