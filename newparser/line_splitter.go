package newparser

func SplitLinesByType(lines []string) (letter string, upper, lower, lyrics []string) {
    types := ClassifyLines(lines)

    for i, typ := range types {
        switch typ {
        case LetterLineType:
            letter = lines[i]
        case UpperAnnotationType:
            upper = append(upper, lines[i])
        case LowerAnnotationType:
            lower = append(lower, lines[i])
        case SyllableType:
            lyrics = append(lyrics, lines[i])
        }
    }

    return
}
