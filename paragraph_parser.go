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

// ✅ LineType Enum
type LineType int

const (
    UpperAnnotationType LineType = iota
    LetterLineType
    LowerAnnotationType
    LyricsType
    UnknownLineType
)

// ✅ DetectLineType determines the type of line based on its contents
func DetectLineType(line string) LineType {
    trimmedLine := strings.TrimSpace(line)

    if len(trimmedLine) > 0 && strings.HasPrefix(trimmedLine, "U:") {
        return UpperAnnotationType
    } else if len(trimmedLine) > 0 && strings.HasPrefix(trimmedLine, "L:") {
        return LyricsType
    } else if len(trimmedLine) > 0 && strings.HasPrefix(trimmedLine, "A:") {
        return LowerAnnotationType
    } else if len(trimmedLine) > 0 {
        return LetterLineType
    }
    return UnknownLineType
}

// ✅ ParseParagraph parses a single paragraph of notation into a Paragraph struct
func ParseParagraph(para string) Paragraph {
    lines := strings.Split(para, "\n")
    var letterLine *LetterLine
    var upperAnnotations []Token
    var lowerAnnotations []Token
    var lyrics []Token

    for _, line := range lines {
        trimmedLine := strings.TrimSpace(line)
        
        if len(trimmedLine) == 0 {
            continue
        }

        // Detect line type and call the appropriate lexer
        switch DetectLineType(trimmedLine) {
        case UpperAnnotationType:
            upperAnnotations = UpperAnnotationLexer(trimmedLine[2:]) // Remove "U:" prefix
        case LetterLineType:
            tokens := LetterLineLexer(trimmedLine)
            letterLine = ParseLetterLine(tokens)
        case LowerAnnotationType:
            lowerAnnotations = LowerAnnotationLexer(trimmedLine[2:]) // Remove "A:" prefix
        case LyricsType:
            lyrics = LyricsLexer(trimmedLine[2:]) // Remove "L:" prefix
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

