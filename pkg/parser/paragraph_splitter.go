package parser

type SplitLines struct {
	LetterLine       string
	UpperAnnotations []string
	LowerAnnotations []string
	LyricLines       []string
}

func SplitParagraph(lines []string) SplitLines {
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
		case LyricLineType:
			result.LyricLines = append(result.LyricLines, lines[i])
		}
	}

	return result
}
