package newparser 

func LexLetterLine(input string) []Token {
    var tokens []Token

    for i, char := range input {
        switch char {
        case 'S', 'r', 'R', 'g', 'G', 'm', 'M', 'P', 'd', 'D', 'n', 'N':
            tokens = append(tokens, Token{Type: Pitch, Value: string(char), Column: i})
        case '-':
            tokens = append(tokens, Token{Type: Dash, Value: string(char), Column: i})
        case '|':
            tokens = append(tokens, Token{Type: Barline, Value: string(char), Column: i})
        case ',':
            tokens = append(tokens, Token{Type: Breath, Value: string(char), Column: i})
        case '(':
            tokens = append(tokens, Token{Type: LeftSlur, Value: string(char), Column: i})
        case ')':
            tokens = append(tokens, Token{Type: RightSlur, Value: string(char), Column: i})
        default:
            tokens = append(tokens, Token{Type: Unknown, Value: string(char), Column: i})
        }
    }

    return tokens
}
