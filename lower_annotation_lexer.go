package main

import (
)

func LowerAnnotationLexer(line string) []Token {
	  Log("DEBUG"," LowerAnnotationLexer line = %s", line)
    var tokens []Token

    for i, char := range line {
        switch char {
        case '.':
            token := Token{
                Type:   LowerOctave,
                Value:  string(char),
                Column: i, // ✅ Exact physical position
            }
            Log("DEBUG", "LowerAnnotationLexer Generated token: Type=%s, Value=%s, Column=%d", token.Type, token.Value, token.Column)
            tokens = append(tokens, token)
        case ':':
            token := Token{
                Type:   LowestOctave,
                Value:  string(char),
                Column: i, // ✅ Exact physical position
            }
            Log("DEBUG", "Generated token: Type=%s, Value=%s, Column=%d", token.Type, token.Value, token.Column)
            tokens = append(tokens, token)
        }
    }
    return tokens
}

