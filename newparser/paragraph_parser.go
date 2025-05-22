package newparser

import (
    "strings"
)

type Paragraph struct {
    Raw              string
    LetterLine       *LetterLine
    UpperAnnotations []string
    LowerAnnotations []string
    LyricLines           []string
}

func ParseParagraph(lines []string) *Paragraph {
    if len(lines) == 0 {
        return nil
    }

    Log("DEBUG", "ParseParagraph raw lines:\n%s", strings.Join(lines, "\n"))

		split := SplitLinesByType(lines)
letter := split.LetterLine
upperLines := split.UpperAnnotations
lowerLines := split.LowerAnnotations
syllableLines := split.LyricLines
    if letter == "" {
        Log("DEBUG", "ParseParagraph aborted: no letter line found.")
        return nil
    }

    Log("DEBUG", "Upper lines: %v", upperLines)
    Log("DEBUG", "Lower lines: %v", lowerLines)
    Log("DEBUG", "Syllable lines: %v", syllableLines)

    tokens := LexLetterLine(letter)
    Log("DEBUG", "Lexed %d tokens from letter line", len(tokens))
    letterLine := ParseLetterLine(letter, tokens)

    var annotations []Annotation
    if len(upperLines) > 0 {
        upper := LexAnnotationLine(upperLines[0], UpperLine)
        Log("DEBUG", "Lexed %d upper annotations", len(upper))
        annotations = append(annotations, upper...)
    }
    if len(lowerLines) > 0 {
        lower := LexAnnotationLine(lowerLines[0], LowerLine)
        Log("DEBUG", "Lexed %d lower annotations", len(lower))
        annotations = append(annotations, lower...)
    }
    if len(syllableLines) > 0 {
        syllables := LexAnnotationLine(syllableLines[0], SyllableLine)
        Log("DEBUG", "Lexed %d syllable annotations", len(syllables))
        annotations = append(annotations, syllables...)
    }

    paragraph := &Paragraph{
        Raw:              strings.Join(lines, "\n"),
        LetterLine:       letterLine,
        UpperAnnotations: upperLines,
        LowerAnnotations: lowerLines,
        LyricLines:           syllableLines,
    }

    Log("DEBUG", "Calling FoldAnnotations with %d annotations", len(annotations))
    FoldAnnotations(paragraph, annotations)

    return paragraph
}
