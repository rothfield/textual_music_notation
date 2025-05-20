
package main

func LetterLineLexer(line string) []Token {
    var tokens []Token

    for i, char := range line {
        switch char {
        case 'S', 'r', 'R', 'g', 'G', 'm', 'M', 'P', 'd', 'D', 'n', 'N':
            tokens = append(tokens, Token{
                Type:   Pitch,
                Value:  string(char),
                Column: i, // ✅ Exact physical position
            })
        case '-':
            tokens = append(tokens, Token{
                Type:   Dash,
                Value:  string(char),
                Column: i, // ✅ Exact physical position
            })
        case '|':
            tokens = append(tokens, Token{
                Type:   Barline,
                Value:  string(char),
                Column: i, // ✅ Exact physical position
            })
        }
    }
    return tokens
}
