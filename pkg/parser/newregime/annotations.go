package parser

type AnnotationRole int

const (
	Upper AnnotationRole = iota
	Lower
	Lyrics
)

type Annotation struct {
	Type   TokenType
	Value  string
	Column int
}

func LexAnnotation(line string, role AnnotationRole) []Annotation {
	switch role {
	case Upper:
		return LexUpperAnnotationLine(line)
	case Lower:
		return LexLowerAnnotationLine(line)
	case Lyrics:
		return LexLyricsAnnotationLine(line)
	default:
		return nil
	}
}
