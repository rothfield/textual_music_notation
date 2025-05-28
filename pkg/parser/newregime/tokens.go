package parser

import (
	"fmt"
)

type TokenType string

const (
	Pitch         TokenType = "Pitch"
	Dash          TokenType = "Dash"
	LeftSlur      TokenType = "LeftSlur"
	RightSlur     TokenType = "RightSlur"
	Slur          TokenType = "Slur"
	Barline       TokenType = "Barline"
	Breath        TokenType = "Breath"
	Octave        TokenType = "Octave"
	Mordent       TokenType = "Mordent"
	Tala          TokenType = "Tala"
	Space         TokenType = "Space"
	LowerOctave   TokenType = "LowerOctave"
	LowestOctave  TokenType = "LowestOctave"
	UpperOctave   TokenType = "UpperOctave"
	HighestOctave TokenType = "HighestOctave"
	Syllable      TokenType = "Syllable"
	Unknown       TokenType = "Unknown"
)

type Token struct {
	Type   TokenType
	Value  string
	Column int
}

func (t TokenType) String() string {
	switch t {
	case Pitch:
		return "pitch"
	case Dash:
		return "dash"
	case Barline:
		return "barline"
	case LeftSlur:
		return "left-slur"
	case RightSlur:
		return "right-slur"
	case Tala:
		return "tala"
	case Syllable:
		return "syllable"
	default:
		return "unknown"
	}
}

func (t Token) String() string {
	return fmt.Sprintf("{%s %s %d}", t.Type, t.Value, t.Column)
}
