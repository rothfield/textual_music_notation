package main

import (
    "strings"
)

type Paragraph struct {
    LetterLine *LetterLine
}

type Composition struct {
    Paragraphs []Paragraph
}

func ParseComposition(input string) *Composition {
    Log("DEBUG", "ParseComposition")
    paragraphs := strings.Split(input, "\n\n")
    var parsedParagraphs []Paragraph

    for _, para := range paragraphs {
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

        parsedParagraphs = append(parsedParagraphs, Paragraph{
            LetterLine: letterLine,
        })
    }

    return &Composition{
        Paragraphs: parsedParagraphs,
    }
}

