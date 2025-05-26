package parser

import (
	"strings"
)

type Composition struct {
	Paragraphs []*Paragraph
	RawText    string
}

func ParseComposition(input string) *Composition {
	Log("DEBUG", "ParseComposition")

	rawText := input
	var lines []string
	lines = strings.Split(input, "\n")

	start := 0
	end := len(lines) - 1

	for start < len(lines) && strings.TrimSpace(lines[start]) == "" {
		start++
	}
	for end >= 0 && strings.TrimSpace(lines[end]) == "" {
		end--
	}

	if start > end {
		return &Composition{}
	}

	trimmedLines := lines[start : end+1]
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

	rawParagraphs := strings.Split(strings.Join(foldedLines, "\n"), "\n\n")
	var parsedParagraphs []*Paragraph

	for _, para := range rawParagraphs {
		if strings.TrimSpace(para) != "" {
			var paragraph *Paragraph
			paragraph = ParseParagraph(strings.Split(para, "\n"))
			if paragraph != nil && paragraph.LetterLine != nil && len(paragraph.LetterLine.Elements) > 0 {
				parsedParagraphs = append(parsedParagraphs, paragraph)
			}
		}
	}

	return &Composition{
		Paragraphs: parsedParagraphs,
		RawText:    rawText,
	}
}
