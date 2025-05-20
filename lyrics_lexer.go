package main

import (
    "strings"
)

func LyricsLexer(line string) []Token {
    tokens := []Token{}
    words := strings.Fields(line)
    currentIndex := 0

    for _, word := range words {
        // Find the exact column position
        columnIndex := strings.Index(line[currentIndex:], word) + currentIndex
        
        // Create the token with the column set
        tokens = append(tokens, Token{
            Type:   Lyrics,
            Value:  word,
            Column: columnIndex,
        })
        
        // Move the current index forward
        currentIndex = columnIndex + len(word)
    }
    return tokens
}

