package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// ComputeMockPath generates a default output path for a mock file.
func ComputeMockPath(searchDir, outputDir, ifacePath, ifaceName string) string {
	// Get relative path from searchDir
	relPath, _ := filepath.Rel(searchDir, filepath.Dir(ifacePath))

	mockSubPath := ""
	for _, s := range strings.Split(relPath, "/") {
		// Add _mocks to the path
		mockSubPath = filepath.Join(mockSubPath, s+"_mocks")
	}

	mockDir := filepath.Join(outputDir, mockSubPath)

	// Convert ifaceName from PascalCase to snake_case
	snakeCaseIfaceName := toSnakeCase(ifaceName)
	snakeCaseIfaceName = strings.ToLower(snakeCaseIfaceName)
	snakeCaseIfaceName = strings.ReplaceAll(snakeCaseIfaceName, "_interface", "")

	// Define mock filename
	mockFile := fmt.Sprintf("%s_mock.go", snakeCaseIfaceName)

	return filepath.Join(mockDir, mockFile)
}

// toSnakeCase converts a PascalCase string to snake_case
func toSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if unicode.IsUpper(r) {
			// Check if the next character is also uppercase
			if i > 0 && (i+1 < len(str) && unicode.IsUpper(rune(str[i+1]))) {
				result = append(result, unicode.ToLower(r))
			} else {
				if i > 0 {
					result = append(result, '_')
				}
				result = append(result, unicode.ToLower(r))
			}
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	return nil
}
