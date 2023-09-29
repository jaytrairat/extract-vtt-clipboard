package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
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

func startExtraction() {
	extractedText, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Cannot read clipboard.")
		return
	}
	preparedText := extractKeyAndValue(extractedText)
	clipboard.WriteAll(preparedText)
	fmt.Printf("Extracted to clipboard at :: %s\n", time.Now().Format("2006-01-02 15:04:05"))
}

var rootCmd = &cobra.Command{
	Use:   "extract-vtt-clipboard",
	Short: "Extract string from Virustotal",
	Run: func(cmd *cobra.Command, args []string) {
		startExtraction()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
