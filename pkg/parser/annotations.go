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
		return LexUpper(line)
	case Lower:
		return LexLower(line)
	case Lyrics:
		return LexLyrics(line)
	default:
		return nil
	}
}
