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

// ✅ ParseParagraph integrates ClassifyLines and lexes only once
func ParseParagraph(para string) Paragraph {
    lines := strings.Split(para, "\n")
    classifiedTypes := ClassifyLines(lines)

    var letterLine *LetterLine
    var upperAnnotations []Token
    var lowerAnnotations []Token
    var lyrics []Token
    lyricsFound := false

    // ✅ Traverse the classified lines in order
    for index, lineType := range classifiedTypes {
        trimmedLine := strings.TrimRight(lines[index], "\n")
        
        switch lineType {
        case UpperAnnotationType:
            upperAnnotations = append(upperAnnotations, UpperAnnotationLexer(trimmedLine)...)
        case LetterLineType:
            tokens := LetterLineLexer(trimmedLine)
            letterLine = ParseLetterLine(tokens)
        case LowerAnnotationType:
            lowerAnnotations = append(lowerAnnotations, LowerAnnotationLexer(trimmedLine)...)
        case LyricsType:
            if !lyricsFound {
                lyrics = append(lyrics, LyricsLexer(trimmedLine)...)
                lyricsFound = true
            } else {
                Log("WARN", "Multiple lyrics lines detected. Ignoring additional ones.")
            }
        default:
            Log("WARN", "Unknown line type detected: %s", trimmedLine)
        }
    }

    // ✅ Use the new constructor to build the Paragraph
    return NewParagraph(letterLine, upperAnnotations, lowerAnnotations, lyrics)
}

