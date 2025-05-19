package main

import (
    "fmt"  // ✅ Import the fmt package
)

func UpperAnnotationLexer(line string) []Token {
    tokens := []Token{}
    for _, char := range line {
        switch char {
        case '0', '+', '1', '2', '3', '4', '5', '6', '7', '8':
            tokens = append(tokens, Token{Type: Tala, Value: string(char)})
        case '~':
            tokens = append(tokens, Token{Type: Mordent, Value: string(char)})
        case '.':
            tokens = append(tokens, Token{Type: HigherOctave, Value: "."})
        case ':':
            tokens = append(tokens, Token{Type: HighestOctave, Value: ":"})
        default:
            Log("WARN", fmt.Sprintf("Unrecognized upper annotation character: %s", string(char)))
        }
    }
    return tokens
}

