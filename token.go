package main

// TokenType represents different types of tokens in the notation.
type TokenType string

const (
    Pitch          TokenType = "Pitch"
    Dash           TokenType = "Dash"
    LeftSlur       TokenType = "LeftSlur"
    RightSlur      TokenType = "RightSlur"
    Slur           TokenType = "Slur"
    Barline        TokenType = "Barline"
    Breath         TokenType = "Breath"
    Lyrics         TokenType = "Lyrics"
    Octave         TokenType = "Octave"
    Mordent        TokenType = "Mordent"
    Tala           TokenType = "Tala"
    Space          TokenType = "Space"
    LowerOctave    TokenType = "LowerOctave"    // ✅ Added
    LowestOctave   TokenType = "LowestOctave"   // ✅ Added
    HigherOctave    TokenType = "HigherOctave"    // ✅ Added
    HighestOctave   TokenType = "HighestOctave"   // ✅ Added
)


// Token represents a single unit of the notation (e.g., a pitch or a dash).
type Token struct {
    Type  TokenType
    Value string
}
    
