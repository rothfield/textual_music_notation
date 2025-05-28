package parser

func LexLine(line string) []Token {
	switch GuessNotation(line) {
	case Western:
		return LexABC(line)
	case Number:
		return Lex123(line)
	default:
		return LexLineSargam(line)
	}
}
