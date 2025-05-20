package main

import (
    "strings"
)

// ✅ LineType Enum
type LineType int

const (
    UpperAnnotationType LineType = iota   // 0
    LetterLineType                        // 1
    LowerAnnotationType                   // 2
    LyricsType                            // 3
    UnknownLineType                       // 4
)


func findLetterLine(lines []string) int {
    maxTokens := 0
    letterLineIndex := -1

    for i, line := range lines {
        tokens := LetterLineLexer(line)

        // ✅ Ignore lines that are just `.`, `:`, or `~`
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

    // ✅ Step 1: Remove leading and trailing blank lines
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

    // ✅ Step 2: First loop to find the LetterLine
    letterLineIndex := findLetterLine(lines)

    if letterLineIndex == -1 {
        Log("ERROR", "No valid LetterLine found during classification.")
        return []LineType{}
    }

    // ✅ Mark the LetterLine — NOTHING ELSE
    types[letterLineIndex] = LetterLineType

    // ✅ Step 3: Second Loop — Start from 0 again
    foundLowerOrLyrics := false
    foundLyrics := false

    for i := 0; i < len(lines); i++ {
			  Log("DEBUG","for loop step 3;  i=%d",i)
				Log("DEBUG", "Types=%s",types)
        if i == letterLineIndex {
            continue // ✅ Skip processing the LetterLine
        }

        // ✅ Upper Annotations are everything before the LetterLine
        if i < letterLineIndex {
            if strings.ContainsAny(lines[i], ".:~") {
                types[i] = UpperAnnotationType
            }
        } else {
            // ✅ After the LetterLine, we look for Lower or Lyrics
            if strings.TrimSpace(lines[i]) != "" {
                if !foundLowerOrLyrics {
                    if strings.ContainsAny(lines[i], ".:~") {
                        types[i] = LowerAnnotationType
                    } else {
                        types[i] = LyricsType
                    }
                    foundLowerOrLyrics = true
                } else if !foundLyrics {
                    types[i] = LyricsType
                    foundLyrics = true
                } else {
                    Log("WARN", "Ignoring additional line at index %d", i)
                }
            }
        }
    }

    // ✅ Log the final classification
    for index, lineType := range types {
        Log("DEBUG", "Line %d classified as %d", index, lineType)
    }

    Log("DEBUG", "ClassifyLines result: %v", types)
    return types
}

func zClassifyLines(lines []string) []LineType {
    Log("DEBUG", "ClassifyLines, lines=%s", lines)

    // ✅ Step 1: Remove leading and trailing blank lines
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

    // ✅ Step 2: First loop to find the LetterLine
    letterLineIndex := findLetterLine(lines)

    if letterLineIndex == -1 {
        Log("ERROR", "No valid LetterLine found during classification.")
        return []LineType{}
    }

    // ✅ Mark the LetterLine
    types[letterLineIndex] = LetterLineType

    // ✅ Step 3: Second Loop — Start from 0 again
    foundLowerOrLyrics := false
    foundLyrics := false

    for i := 0; i < len(lines); i++ {
        if i == letterLineIndex {
            continue // ✅ Skip processing the LetterLine
        }

        if i < letterLineIndex {
            // ✅ Explicitly check if it is valid for UpperAnnotation
            if strings.ContainsAny(lines[i], ".:~") {
                types[i] = UpperAnnotationType
            } else {
                types[i] = UnknownLineType
            }
        } else if i > letterLineIndex {
            if strings.TrimSpace(lines[i]) != "" {
                if !foundLowerOrLyrics {
                    if strings.ContainsAny(lines[i], ".:~") {
                        types[i] = LowerAnnotationType
                    } else {
                        types[i] = LyricsType
                    }
                    foundLowerOrLyrics = true
                } else if !foundLyrics {
                    types[i] = LyricsType
                    foundLyrics = true
                } else {
                    Log("WARN", "Ignoring additional line at index %d", i)
                }
            }
        }
    }

    Log("DEBUG", "ClassifyLines result: %v", types)
    return types
}


