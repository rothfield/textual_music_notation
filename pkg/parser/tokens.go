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
	Barline       TokenType = "Barline"
	Breath        TokenType = "Breath"
	Octave        TokenType = "Octave"
	Mordent       TokenType = "Mordent"
	Tala          TokenType = "Tala"
	LowerOctave   TokenType = "LowerOctave"
	LowestOctave  TokenType = "LowestOctave"
	UpperOctave   TokenType = "UpperOctave"
	HighestOctave TokenType = "HighestOctave"
	Syllable      TokenType = "Syllable"
	Space         TokenType = "Space"
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
	case Space:
		return "space" // <-- ADD THIS LINE

	default:
		return "unknown"
	}
}

func (t Token) String() string {
	return fmt.Sprintf("{%s %s %d}", t.Type, t.Value, t.Column)
}
