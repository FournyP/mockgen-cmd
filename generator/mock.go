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

	cmd := exec.Command(
		"mockgen",
		"-source="+interfacePath,
		"-destination="+outputPath,
		"-package=mocks",
	)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("mockgen execution failed: %w", err)
	}
	return nil
}
