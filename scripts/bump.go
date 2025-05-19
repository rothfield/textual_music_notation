package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

// Bump the version in the file URLs
func bumpVersion(htmlContent string) string {
	// Regular expressions to match the version query parameter for main.js and style.css
	jsRegex := regexp.MustCompile(`(main\.js\?v=)(\d+)`)
	cssRegex := regexp.MustCompile(`(style\.css\?v=)(\d+)`)

	// Bump the version number (increment by 1)
	htmlContent = jsRegex.ReplaceAllStringFunc(htmlContent, func(s string) string {
		parts := strings.Split(s, "?v=")
		version := parts[1]
		newVersion := fmt.Sprintf("%d", atoi(version)+1)
		return fmt.Sprintf("%s?v=%s", parts[0], newVersion)
	})

	htmlContent = cssRegex.ReplaceAllStringFunc(htmlContent, func(s string) string {
		parts := strings.Split(s, "?v=")
		version := parts[1]
		newVersion := fmt.Sprintf("%d", atoi(version)+1)
		return fmt.Sprintf("%s?v=%s", parts[0], newVersion)
	})

	return htmlContent
}

// Helper to convert string to integer
func atoi(s string) int {
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
	}
	return result
}

func main() {
	// Read the contents of the index.html file
	htmlFilePath := "index.html"
	content, err := ioutil.ReadFile(htmlFilePath)
	if err != nil {
		log.Fatalf("Failed to read index.html: %v", err)
	}

	// Update the version numbers in the content
	updatedContent := bumpVersion(string(content))

	// Write the updated content back to index.html
	err = ioutil.WriteFile(htmlFilePath, []byte(updatedContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write updated content to index.html: %v", err)
	}

	log.Println("index.html has been updated with new versions.")
}

