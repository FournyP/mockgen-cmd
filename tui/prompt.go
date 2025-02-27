package tui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PromptInput asks the user for input
func PromptInput(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question + " ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// PromptYesNoWithDefaultValue asks a yes/no question with a default answer
func PromptYesNoWithDefaultValue(question string, defaultValue bool) bool {
	defaultStr := "n"
	if defaultValue {
		defaultStr = "y"
	}

	for {
		fmt.Printf("%s (y/n) [%s]: ", question, defaultStr)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "" {
			return defaultValue // Use default if empty
		}
		if input == "y" || input == "yes" {
			return true
		}
		if input == "n" || input == "no" {
			return false
		}
		fmt.Println("Invalid input. Please enter 'y' or 'n'.")
	}
}

// PromptInputWithDefault asks for input but provides a default value
func PromptInputWithDefault(question, defaultValue string) string {
	fmt.Printf("%s [%s]: ", question, defaultValue)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return defaultValue
	}
	return input
}
