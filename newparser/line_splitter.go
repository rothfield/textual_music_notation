package newparser

type SplitLines struct {
    LetterLine       string
    UpperAnnotations []string
    LowerAnnotations []string
    Syllables           []string
}

func SplitLinesByType(lines []string) SplitLines {
    types := ClassifyLines(lines)

    var result SplitLines

    for i, typ := range types {
        switch typ {
        case LetterLineType:
            result.LetterLine = lines[i]
        case UpperAnnotationType:
            result.UpperAnnotations = append(result.UpperAnnotations, lines[i])
        case LowerAnnotationType:
            result.LowerAnnotations = append(result.LowerAnnotations, lines[i])
        case SyllableType:
            result.Syllables = append(result.Syllables, lines[i])
        }
    }

    return result
}
