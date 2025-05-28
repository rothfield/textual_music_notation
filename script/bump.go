package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

// ✅ Bump the version in the file URLs using current timestamp
func bumpVersion(htmlContent string) string {
	log.Println("🔄 Replacing version with timestamp...")

	// ✅ Match any <script src="...js?v=..."> and <link href="...css?v=...">
	jsRegex := regexp.MustCompile(`(<script\s+[^>]*src=")([^"]*\.js\?v=)(\d+)(".*?>\s*</script>)`)
	cssRegex := regexp.MustCompile(`(<link\s+[^>]*href=")([^"]*\.css\?v=)(\d+)(".*?>)`)

	// ✅ Use current UNIX timestamp
	timestamp := fmt.Sprintf("%d", time.Now().Unix())

	// ✅ Replace JS version
	htmlContent = jsRegex.ReplaceAllString(htmlContent, fmt.Sprintf("${1}${2}%s${4}", timestamp))

	// ✅ Replace CSS version
	htmlContent = cssRegex.ReplaceAllString(htmlContent, fmt.Sprintf("${1}${2}%s${4}", timestamp))

	log.Printf("✅ Updated all versions to timestamp: %s", timestamp)
	return htmlContent
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Current working directory:", dir)

	// ✅ Path to index.html (adjust if needed)
	htmlFilePath := "./web/index.html"

	// ✅ Read the contents of the index.html file
	content, err := os.ReadFile(htmlFilePath)
	if err != nil {
		log.Fatalf("Failed to read index.html: %v", err)
	}

	// ✅ Update the version numbers in the content
	updatedContent := bumpVersion(string(content))

	// ✅ Write the updated content back to index.html
	err = os.WriteFile(htmlFilePath, []byte(updatedContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write updated content to index.html: %v", err)
	}

	log.Println("✅ index.html has been updated with timestamped version numbers.")
}

