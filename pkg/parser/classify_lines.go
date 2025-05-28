package parser

import (
	"strings"
	"textual_music_notation/internal/logger"
)

func findPitchLine(lines []string) int {
	logger.Log("DEBUG", "findPitchLine %s", lines)
	maxTokens := 0
	bestIndex := -1

	for i, line := range lines {
		tokens := LexLine(line)
		Log("DEBUG", "findPitchLine: line %d => %d tokens", i, len(tokens))
		Log("DEBUG", "findPitchLine: tokens %s ", tokens)

		if strings.TrimSpace(line) == "." || strings.TrimSpace(line) == ":" || strings.TrimSpace(line) == "~" {
			continue
		}

		unknowns := 0
		for _, tok := range tokens {
			if tok.Type == Unknown {
				unknowns++
			}
		}

		if unknowns > 1 {
			continue
		}

		if len(tokens) > maxTokens {
			maxTokens = len(tokens)
			bestIndex = i
		}
	}

	if bestIndex == -1 {
		Log("WARN", "findPitchLine: no valid line found, falling back to longest tokenizable line")
		maxTokens = 0
		for i, line := range lines {
			tokens := LexLine(line)
			if len(tokens) > maxTokens {
				maxTokens = len(tokens)
				bestIndex = i
			}
		}
	}

	Log("DEBUG", "Letter line identified at index %d", bestIndex)
	return bestIndex
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
	for i := range types {
		types[i] = UnknownLineType
	}

	letterLineIndex := findPitchLine(lines)
	if letterLineIndex == -1 {
		Log("ERROR", "No valid PitchLine found during classification.")
		return []LineType{}
	}
	types[letterLineIndex] = PitchLineType

	foundLowerOrLyric := false
	foundLyric := false

	for i := 0; i < len(lines); i++ {
		Log("DEBUG", "for loop step 3;  i=%d", i)
		Log("DEBUG", "Types=%s", types)
		if i == letterLineIndex {
			continue
		}

		if types[i] != UnknownLineType {
			continue
		}

		if i < letterLineIndex {
			if strings.ContainsAny(lines[i], ".:~") {
				types[i] = UpperAnnotationType
			}
		} else {
			if strings.TrimSpace(lines[i]) != "" {
				if !foundLowerOrLyric {
					if strings.ContainsAny(lines[i], ".:~") {
						types[i] = LowerAnnotationType
					} else {
						types[i] = LyricLineType
					}
					foundLowerOrLyric = true
				} else if !foundLyric {
					types[i] = LyricLineType
					foundLyric = true
				} else {
					Log("WARN", "Ignoring additional line at index %d", i)
				}
			}
		}
	}

	for index, lineType := range types {
		Log("DEBUG", "Line %d classified as %s", index, LineTypeNames[lineType])
	}

	Log("DEBUG", "ClassifyLines result: %s", types)
	return types
}
