package newparser

type SplitLines struct {
    LetterLine       string
    UpperAnnotations []string
    LowerAnnotations []string
    LyricLines           []string
}

func SplitLinesByType(lines []string) SplitLines {
    types := ClassifyLines(lines)

    var result SplitLines

    for i, typ := range types {
        switch typ {
        case LetterLineRole:
            result.LetterLine = lines[i]
        case UpperAnnotationsLineRole:
            result.UpperAnnotations = append(result.UpperAnnotations, lines[i])
        case LowerAnnotationsLineRole:
            result.LowerAnnotations = append(result.LowerAnnotations, lines[i])
        case LyricLineRole:
            result.LyricLines = append(result.LyricLines, lines[i])
        }
    }

    return result
}
