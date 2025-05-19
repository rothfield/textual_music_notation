package main

import (
    "strings"
)

func LyricsLexer(line string) []Token {
    tokens := []Token{}
    words := strings.Fields(line)
    for _, word := range words {
        tokens = append(tokens, Token{Type: Lyrics, Value: word})
    }
    return tokens
}

