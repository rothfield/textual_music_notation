package parser

type LetterLine struct {
	Elements []Element
	Raw      string
}

type Element struct {
	Token          Token
	Column         int
	Octave         int
	Mordent        bool
	Tala           string
	Syllable       string
	ExtraSyllables []string
	IsBeat         bool
	SubElements    []Element
	Divisions      int
}

func (t Token) GetColumn() int {
	return t.Column
}

func (t Token) GetType() TokenType {
	return t.Type
}

func (p Element) GetColumn() int {
	return p.Column
}

func (p Element) GetType() TokenType {
	return p.Token.Type
}
