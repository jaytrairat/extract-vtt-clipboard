package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/atotto/clipboard"
)

func extractKeyAndValue(raw string) string {
	var result strings.Builder
	lines := strings.Split(raw, "\n")

	if len(lines)%2 == 0 {
		for i := 0; i < len(lines); i += 2 {
			key := strings.TrimSpace(lines[i])
			value := strings.TrimSpace(lines[i+1])
			result.WriteString(fmt.Sprintf("%s\t%s\n", key, value))
		}
	}
	fmt.Printf("Number of lines is :: %d\n", len(lines)/2)
	return result.String()
}

func main() {
	extractedText, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Cannot read clipboard.")
		return
	}
	preparedText := extractKeyAndValue(extractedText)
	clipboard.WriteAll(preparedText)
	fmt.Printf("Extracted to clipboard at :: %s\n", time.Now().Format("2006-01-02 15:04:05"))
}
