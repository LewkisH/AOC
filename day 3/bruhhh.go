/*package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	fileB, _ := os.ReadFile("day3.txt")
	input := string(fileB)
	result := removeDigitsAndDots(input)

	fmt.Println("String without digits and dots:", result)
}

func removeDigitsAndDots(input string) string {
	// Regular expression to match digits and dots
	re := regexp.MustCompile(`[\d.\n]`)
	// Replace all matches with an empty string
	result := re.ReplaceAllString(input, "")
	return result
}
*/