package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Exists checks if a file or directory exists
func Exists(path string) bool {
	// Use os.Stat to get the file information
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// If the file does not exist, return false
		return false
	}
	// If the file exists, return true
	return true
}

// GetInput gets user input
func GetInput(msg string) string {
	// Print the message to the console
	fmt.Println(msg)
	// Create a new reader for the standard input
	reader := bufio.NewReader(os.Stdin)
	// Read a line of input from the user
	input, err := reader.ReadString('\n')
	if err != nil {
		// If there is an error, return an empty string
		return ""
	}
	// Remove the newline character from the end of the input
	return strings.Replace(strings.TrimSuffix(input, "\n"), "\r", "", -1)
}
