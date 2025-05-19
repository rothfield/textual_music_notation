package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

// ✅ Bump the version in the file URLs
func bumpVersion(htmlContent string) string {
	log.Println("🔄 Bumping versions for all JS and CSS...")

	// ✅ Regular expressions to match any JS and CSS links
	jsRegex := regexp.MustCompile(`(<script\s+.*src=")([^"]*\.js\?v=)(\d+)(".*><\/script>)`)
	cssRegex := regexp.MustCompile(`(<link\s+.*href=")([^"]*\.css\?v=)(\d+)(".*>)`)

	// ✅ Bump JS version
	htmlContent = jsRegex.ReplaceAllStringFunc(htmlContent, func(s string) string {
		parts := strings.Split(s, "?v=")
		version := parts[1][:strings.Index(parts[1], "\"")]
		newVersion := fmt.Sprintf("%d", atoi(version)+1)
		log.Printf("✅ Bumped JS version to %s", newVersion)
		// ✅ Fixed formatting here
		return fmt.Sprintf("%s?v=%s%s", parts[0], newVersion, parts[1][strings.Index(parts[1], "\""):])
	})

	// ✅ Bump CSS version
	htmlContent = cssRegex.ReplaceAllStringFunc(htmlContent, func(s string) string {
		parts := strings.Split(s, "?v=")
		version := parts[1][:strings.Index(parts[1], "\"")]
		newVersion := fmt.Sprintf("%d", atoi(version)+1)
		log.Printf("✅ Bumped CSS version to %s", newVersion)
		// ✅ Fixed formatting here
		return fmt.Sprintf("%s?v=%s%s", parts[0], newVersion, parts[1][strings.Index(parts[1], "\""):])
	})

	return htmlContent
}

// ✅ Helper to convert string to integer
func atoi(s string) int {
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
	}
	return result
}

func main() {
	// ✅ Path to index.html (adjusted to web folder)
	htmlFilePath := "../web/index.html" 

	// ✅ Read the contents of the index.html file
	content, err := ioutil.ReadFile(htmlFilePath)
	if err != nil {
		log.Fatalf("Failed to read index.html: %v", err)
	}

	// ✅ Update the version numbers in the content
	updatedContent := bumpVersion(string(content))

	// ✅ Write the updated content back to index.html
	err = ioutil.WriteFile(htmlFilePath, []byte(updatedContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write updated content to index.html: %v", err)
	}

	log.Println("✅ index.html has been updated with new version numbers.")
}

