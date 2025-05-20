
package main

func UpperAnnotationLexer(line string) []Token {
    var tokens []Token

    for i, char := range line {
        switch char {
        case '.', ':': // Octaves
            tokens = append(tokens, Token{
                Type:   HigherOctave,
                Value:  string(char),
                Column: i, // ✅ Exact physical position
            })
        case '~': // Mordent
            tokens = append(tokens, Token{
                Type:   Mordent,
                Value:  string(char),
                Column: i, // ✅ Exact physical position
            })
        case '0', '+', '1', '2', '3', '4', '5', '6', '7', '8': // Talas
            tokens = append(tokens, Token{
                Type:   Tala,
                Value:  string(char),
                Column: i, // ✅ Exact physical position
            })
        }
    }
    return tokens
}
