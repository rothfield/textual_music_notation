package newparser

type Paragraph struct {
    LetterLine *LetterLine
}

func ParseParagraph(lines []string) *Paragraph {
    if len(lines) == 0 {
        return nil
    }

    letter, upperLines, lowerLines, lyricLines := SplitLinesByType(lines)
    if letter == "" {
        return nil
    }

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
        LetterLine: letterLine,
    }

    FoldAnnotations(paragraph, annotations)

    return paragraph
}
