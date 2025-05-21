package newparser

import (
    "strings"
)
var lineTypeNames = map[LineType]string{
    LetterLineType:      "LetterLine",
    UpperAnnotationType: "UpperAnnotation",
    LowerAnnotationType: "LowerAnnotation",
    SyllableType:        "Syllable",
}

func findLetterLine(lines []string) int {
    maxTokens := 0
    letterLineIndex := -1

    for i, line := range lines {
        tokens := LexLetterLine(line) // Use the existing lexer

        if strings.TrimSpace(line) == "." || strings.TrimSpace(line) == ":" || strings.TrimSpace(line) == "~" {
            continue
        }

        if len(tokens) > maxTokens {
            maxTokens = len(tokens)
            letterLineIndex = i
        }
    }

    return letterLineIndex
}

func ClassifyLines(lines []string) []LineType {
    Log("DEBUG", "ClassifyLines, lines=%s", lines)

    start := 0
    end := len(lines) - 1
    for start < len(lines) && strings.TrimSpace(lines[start]) == "" {
        start++
    }
    for end >= 0 && strings.TrimSpace(lines[end]) == "" {
        end--
    }

    if start > end {
        Log("WARN", "Empty paragraph after trimming blank lines.")
        return []LineType{}
    }

    lines = lines[start : end+1]
    types := make([]LineType, len(lines))

    letterLineIndex := findLetterLine(lines)
    if letterLineIndex == -1 {
        Log("ERROR", "No valid LetterLine found during classification.")
        return []LineType{}
    }
    types[letterLineIndex] = LetterLineType

    foundLowerOrSyllable := false
    foundSyllable := false

    for i := 0; i < len(lines); i++ {
        Log("DEBUG", "for loop step 3;  i=%d", i)
        Log("DEBUG", "Types=%s", types)
        if i == letterLineIndex {
            continue
        }

        if i < letterLineIndex {
            if strings.ContainsAny(lines[i], ".:~") {
                types[i] = UpperAnnotationType
            }
        } else {
            if strings.TrimSpace(lines[i]) != "" {
                if !foundLowerOrSyllable {
                    if strings.ContainsAny(lines[i], ".:~") {
                        types[i] = LowerAnnotationType
                    } else {
                        types[i] = SyllableType
                    }
                    foundLowerOrSyllable = true
                } else if !foundSyllable {
                    types[i] = SyllableType
                    foundSyllable = true
                } else {
                    Log("WARN", "Ignoring additional line at index %d", i)
                }
            }
        }
    }

    for index, lineType := range types {
			Log("DEBUG", "Line %d classified as %s", index, lineTypeNames[lineType])
    }

    Log("DEBUG", "ClassifyLines result: %s", types)
    return types
}
