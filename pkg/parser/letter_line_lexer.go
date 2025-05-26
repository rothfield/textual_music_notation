package parser

func LexLetterLine(line string) []Token {
    switch GuessPitchSystem(line) {
    case Western:
        return LexLetterLineWestern(line)
    case Number:
        return LexLetterLineNumber(line)
    default:
        return LexLetterLineSargam(line)
    }
}

