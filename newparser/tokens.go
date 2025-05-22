package newparser

type TokenType string

const (
    // Structural tokens (renamed to avoid struct conflict)
    PitchToken     TokenType = "Pitch"
    DashToken      TokenType = "Dash"
    BarlineToken   TokenType = "Barline"
    BreathToken    TokenType = "Breath"
    PhraseToken    TokenType = "Phrase"

    // System tokens (unchanged)
    LeftSlur       TokenType = "LeftSlur"
    RightSlur      TokenType = "RightSlur"
    Whitespace     TokenType = "Whitespace"
    Error          TokenType = "Error"
    Unknown        TokenType = "Unknown"

    // Annotation tokens (unchanged)
    HigherOctave   TokenType = "HigherOctave" // :
    LowerOctave    TokenType = "LowerOctave"  // .
    Mordent        TokenType = "Mordent"      // ~
    Tala           TokenType = "Tala"
    TalaMarker     TokenType = "TalaMarker"
    Syllable       TokenType = "Syllable"
    UpperOctave    TokenType = "UpperOctave"
)

type Token struct {
	Type   TokenType
	Value  string
	Column int
}

