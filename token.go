
package main

// TokenType represents different types of tokens in the notation.
type TokenType string

const (
    Pitch   TokenType = "Pitch"
    Dash    TokenType = "Dash"
    Slur    TokenType = "Slur"
    Barline TokenType = "Barline"
    Breath  TokenType = "Breath"
    Lyrics  TokenType = "Lyrics"
    Octave  TokenType = "Octave"
    Mordent TokenType = "Mordent"
    Tala    TokenType = "Tala"
    Space   TokenType = "Space"
)

// Token represents a single unit of the notation (e.g., a pitch or a dash).
type Token struct {
    Type  TokenType
    Value string
}
    