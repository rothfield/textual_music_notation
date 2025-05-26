package parser

type AnnotationRole int

const (
	UpperLine AnnotationRole = iota
	LowerLine
	SyllableLine
)

type Annotation struct {
	Type   TokenType
	Value  string
	Column int
}

func LexAnnotationLine(line string, role AnnotationRole) []Annotation {
	switch role {
	case UpperLine:
		return LexUpperAnnotationLine(line)
	case LowerLine:
		return LexLowerAnnotationLine(line)
	case SyllableLine:
		return LexLyricsAnnotationLine(line)
	default:
		return nil
	}
}
