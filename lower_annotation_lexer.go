package main

import (
    "fmt"  // ✅ Import the fmt package
)

func LowerAnnotationLexer(line string) []Token {
    tokens := []Token{}
    for _, char := range line {
        switch char {
        case '.':
            tokens = append(tokens, Token{Type: LowerOctave, Value: "."})
        case ':':
            tokens = append(tokens, Token{Type: LowestOctave, Value: ":"})
        default:
            Log("WARN", fmt.Sprintf("Unrecognized lower annotation character: %s", string(char)))
        }
    }
    return tokens
}

