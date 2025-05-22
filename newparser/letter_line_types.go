package newparser

type LetterLine struct {
    Elements []LetterLineElement
    Raw      string
}

type LetterLineElement struct {
    Token        Token
    Column       int
    Octave       int
    Mordent      bool
    TalaMarker   string
    SyllableText string
    IsBeat       bool
    SubElements  []LetterLineElement
    Divisions    int
}

func (t Token) GetColumn() int {
    return t.Column
}

func (t Token) GetType() TokenType {
    return t.Type
}

func (p LetterLineElement) GetColumn() int {
    return p.Column
}

func (p LetterLineElement) GetType() TokenType {
    return p.Token.Type
}
