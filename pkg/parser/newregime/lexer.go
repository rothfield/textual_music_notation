package parser

func LexLetterLine(line string) []Token {
	switch GuessNotation(line) {
	case Western:
		return LexABC(line)
	case Number:
		return LexLetterLineNumber(line)
	default:
		return LexLetterLineSargam(line)
	}
}
