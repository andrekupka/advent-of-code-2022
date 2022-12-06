package util

import (
	"os"
)

func ReadFileAsString(path string) (string, error) {
	contentAsBytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(contentAsBytes), nil
}
