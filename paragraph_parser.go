package main

import (
    "strings"
)

// Paragraph represents a single block of notation
type Paragraph struct {
    LetterLine *LetterLine
}

// ParseParagraph parses a single paragraph of notation into a Paragraph struct
func ParseParagraph(para string) Paragraph {
    lines := strings.Split(para, "\n")
    var letterLine *LetterLine

    for _, line := range lines {
        trimmedLine := strings.TrimSpace(line)
        
        if len(trimmedLine) == 0 {
            continue
        }

        if letterLine == nil {
            tokens := Lexer(trimmedLine)
            letterLine = ParseLetterLine(tokens)
        }
    }

    return Paragraph{
        LetterLine: letterLine,
    }
}

