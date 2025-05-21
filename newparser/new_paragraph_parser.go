package newparser

import (
    "strings"
)

type Paragraph struct {
    Raw              string
    LetterLine       *LetterLine
    UpperAnnotations []string
    LowerAnnotations []string
    Lyrics           []string
}

func ParseParagraph(lines []string) *Paragraph {
    if len(lines) == 0 {
        return nil
    }

    Log("DEBUG", "ParseParagraph raw lines:\n%s", strings.Join(lines, "\n"))

    letter, upperLines, lowerLines, lyricLines := SplitLinesByType(lines)
    if letter == "" {
        Log("DEBUG", "ParseParagraph aborted: no letter line found.")
        return nil
    }

    Log("DEBUG", "Upper lines: %v", upperLines)
    Log("DEBUG", "Lower lines: %v", lowerLines)
    Log("DEBUG", "Lyric lines: %v", lyricLines)

    tokens := LexLetterLine(letter)
    letterLine := ParseLetterLine(letter, tokens)

    var annotations []Annotation
    if len(upperLines) > 0 {
        annotations = append(annotations, LexAnnotationLine(upperLines[0], UpperLine)...)
    }
    if len(lowerLines) > 0 {
        annotations = append(annotations, LexAnnotationLine(lowerLines[0], LowerLine)...)
    }
    if len(lyricLines) > 0 {
        annotations = append(annotations, LexAnnotationLine(lyricLines[0], SyllableLine)...)
    }

    paragraph := &Paragraph{
        Raw:              strings.Join(lines, "\n"),
        LetterLine:       letterLine,
        UpperAnnotations: upperLines,
        LowerAnnotations: lowerLines,
        Lyrics:           lyricLines,
    }

    Log("DEBUG", "Calling FoldAnnotations with %d annotations", len(annotations))
    FoldAnnotations(paragraph, annotations)

    return paragraph
}

