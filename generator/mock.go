package generator

import (
	"fmt"
	"os/exec"
	"strings"
)

// GenerateMock runs mockgen to generate a mock for an interface.
func GenerateMock(interfaceName, interfacePath, outputPath string) error {
	// Get the output directory
	splitedOutputPath := strings.Split(outputPath, "/")
	outputDir := strings.Join(splitedOutputPath[:len(splitedOutputPath)-1], "/")

	if err := CreateDirIfNotExist(outputDir); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	var packageName string

	if len(splitedOutputPath) > 1 {
		packageName = splitedOutputPath[len(splitedOutputPath)-2]
	} else {
		packageName = "mocks"
	}

	cmd := exec.Command(
		"mockgen",
		"-source="+interfacePath,
		"-destination="+outputPath,
		"-package="+packageName,
	)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("mockgen execution failed: %w", err)
	}
	return nil
}
