package main

import (
    "strings"
    "fmt"
)

// ✅ LineType Enum
type LineType int

const (
    UpperAnnotationType LineType = iota
    LetterLineType
    LowerAnnotationType
    LyricsType
    UnknownLineType
)

// ✅ ClassifyLines determines the type of each line in the paragraph
func ClassifyLines(lines []string) map[string]LineType {
    lineTypes := make(map[string]LineType)

    // ✅ Step 1: Remove leading and trailing blank lines
    start := 0
    end := len(lines) - 1

    // Skip leading blanks
    for start < len(lines) && strings.TrimSpace(lines[start]) == "" {
        start++
    }

    // Skip trailing blanks
    for end >= 0 && strings.TrimSpace(lines[end]) == "" {
        end--
    }

    // If nothing remains, return empty
    if start > end {
        Log("WARN", "Empty paragraph after trimming blank lines.")
        return lineTypes
    }

    // ✅ Create a new trimmed slice of lines
    lines = lines[start : end+1]

    // ✅ If there is only one line, it is automatically the LetterLine
    if len(lines) == 1 {
        trimmedLine := strings.TrimSpace(lines[0])
        if len(trimmedLine) > 0 {
            lineTypes[trimmedLine] = LetterLineType
            return lineTypes
        }
    }

    // Step 2: Identify the LetterLine (contains "|")
    letterLine := ""
    letterLineIndex := -1

    for i, line := range lines {
        trimmedLine := strings.TrimSpace(line)

        // If it contains |, prioritize this as the letter line
        if strings.Contains(trimmedLine, "|") {
            letterLine = trimmedLine
            letterLineIndex = i
            break
        }
    }

    // ✅ If no LetterLine with "|" is found, take the **longest line** with valid characters
    if letterLine == "" {
        validCharacters := "SrRgGmMdDnNP-"
        for i, line := range lines {
            trimmedLine := strings.TrimSpace(line)

            // Ensure all characters are in the valid list
            if strings.IndexFunc(trimmedLine, func(r rune) bool {
                return !strings.ContainsRune(validCharacters, r)
            }) == -1 {
                if len(trimmedLine) > len(letterLine) {
                    letterLine = trimmedLine
                    letterLineIndex = i
                }
            }
        }
    }

    // ✅ If still no LetterLine is found, just log a warning and continue
    if letterLine == "" {
        Log("WARN", fmt.Sprintf("No valid LetterLine found in paragraph: %v", lines))
        return lineTypes
    }

    // Mark the LetterLine
    lineTypes[letterLine] = LetterLineType

    // Step 3: Classify lines **before** the LetterLine as Upper Annotations
    for i := 0; i < letterLineIndex; i++ {
        trimmedLine := strings.TrimSpace(lines[i])
        if len(trimmedLine) > 0 {
            lineTypes[trimmedLine] = UpperAnnotationType
        }
    }

    // Step 4: Classify lines **after** the LetterLine
    foundLower := false
    for i := letterLineIndex + 1; i < len(lines); i++ {
        trimmedLine := strings.TrimSpace(lines[i])
        if len(trimmedLine) == 0 {
            continue
        }

        // The first non-empty line after LetterLine is LowerAnnotation
        if !foundLower {
            lineTypes[trimmedLine] = LowerAnnotationType
            foundLower = true
        } else {
            lineTypes[trimmedLine] = LyricsType
        }
    }

    return lineTypes
}

