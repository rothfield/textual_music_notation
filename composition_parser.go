package main

import (
    "strings"
)

// ✅ Composition represents the entire parsed structure
type Composition struct {
    Paragraphs []Paragraph
}

// ✅ ParseComposition parses the input notation into a Composition structure
func ParseComposition(input string) *Composition {
    Log("DEBUG", "ParseComposition")

    // ✅ Step 1: Split by newline and trim leading/trailing blanks
    lines := strings.Split(input, "\n")

    start := 0
    end := len(lines) - 1

    // Trim leading blanks
    for start < len(lines) && strings.TrimSpace(lines[start]) == "" {
        start++
    }

    // Trim trailing blanks
    for end >= 0 && strings.TrimSpace(lines[end]) == "" {
        end--
    }

    if start > end {
        // If all lines are blank, return an empty composition
        return &Composition{}
    }

    // ✅ Step 2: Slice the lines and fold consecutive empty lines
    trimmedLines := lines[start : end+1]

    // ✅ Step 3: Fold consecutive empty lines
    foldedLines := []string{}
    consecutiveBlank := false

    for _, line := range trimmedLines {
        if strings.TrimSpace(line) == "" {
            if !consecutiveBlank {
                foldedLines = append(foldedLines, "")
                consecutiveBlank = true
            }
        } else {
            foldedLines = append(foldedLines, line)
            consecutiveBlank = false
        }
    }

    // ✅ Step 4: Split by folded blank lines to get paragraphs
    rawParagraphs := strings.Split(strings.Join(foldedLines, "\n"), "\n\n")

    // ✅ Step 5: Parse each paragraph separately, ignoring empty ones
    var parsedParagraphs []Paragraph
    for _, para := range rawParagraphs {
        if strings.TrimSpace(para) != "" {
            paragraph := ParseParagraph(para)
            if paragraph.LetterLine != nil && len(paragraph.LetterLine.Elements) > 0 {
                parsedParagraphs = append(parsedParagraphs, paragraph)
            }
        }
    }

    // ✅ Step 6: Return the full composition structure
    return &Composition{
        Paragraphs: parsedParagraphs,
    }
}

