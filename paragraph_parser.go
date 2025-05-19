package main

import (
    "strings"
)

// ✅ Paragraph represents a single block of notation
type Paragraph struct {
    LetterLine        *LetterLine
    UpperAnnotations  []Token
    LowerAnnotations  []Token
    Lyrics            []Token
}

// ✅ ParseParagraph parses a single paragraph of notation into a Paragraph struct
func ParseParagraph(para string) Paragraph {
    lines := strings.Split(para, "\n")
    lineTypes := ClassifyLines(lines)

    var letterLine *LetterLine
    var upperAnnotations []Token
    var lowerAnnotations []Token
    var lyrics []Token

    for _, line := range lines {
        trimmedLine := strings.TrimSpace(line)
        switch lineTypes[trimmedLine] {
        case UpperAnnotationType:
            upperAnnotations = UpperAnnotationLexer(trimmedLine)
        case LetterLineType:
            tokens := LetterLineLexer(trimmedLine)
            letterLine = ParseLetterLine(tokens)
        case LowerAnnotationType:
            lowerAnnotations = LowerAnnotationLexer(trimmedLine)
        case LyricsType:
            lyrics = LyricsLexer(trimmedLine)
        default:
            Log("WARN", "Unknown line type detected")
        }
    }

    return Paragraph{
        LetterLine:        letterLine,
        UpperAnnotations:  upperAnnotations,
        LowerAnnotations:  lowerAnnotations,
        Lyrics:            lyrics,
    }
}

