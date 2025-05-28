package parser

func LexLine(line string) []Token {
	switch GuessNotation(line) {
	case Western:
		return LexABC(line)
	case Number:
		return LexLineNumber(line)
	default:
		return LexLineSargam(line)
	}
}
