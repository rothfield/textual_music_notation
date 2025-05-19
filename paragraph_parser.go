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

// ✅ NewParagraph constructs a Paragraph, applies folding, and returns it.
func NewParagraph(letterLine *LetterLine, upperAnnotations []Token, lowerAnnotations []Token, lyrics []Token) Paragraph {
    paragraph := Paragraph{
        LetterLine:        letterLine,
        UpperAnnotations:  upperAnnotations,
        LowerAnnotations:  lowerAnnotations,
        Lyrics:            lyrics,
    }
    
    // Apply the folding phase to map annotations to the LetterLine elements.
    FoldAnnotations(&paragraph)

    return paragraph
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

    // ✅ Use the new constructor
    return NewParagraph(letterLine, upperAnnotations, lowerAnnotations, lyrics)
}

