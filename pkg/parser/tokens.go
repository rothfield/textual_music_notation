package parser

import (
	"fmt"
)

//go:generate stringer -type=TokenType -trimprefix=TokenType

type TokenType int

const (
	TokenTypeUnknown TokenType = iota
	TokenTypePitch
	TokenTypeDash
	TokenTypeLeftSlur
	TokenTypeRightSlur
	TokenTypeBarline
	TokenTypeBreath
	TokenTypeOctave
	TokenTypeMordent
	TokenTypeTala
	TokenTypeLowerOctave
	TokenTypeLowestOctave
	TokenTypeUpperOctave
	TokenTypeHighestOctave
	TokenTypeSyllable
	TokenTypeSpace
)

type Token struct {
	Type   TokenType
	Value  string
	Column int
}

func (t Token) String() string {
	return fmt.Sprintf("{%s %s %d}", t.Type, t.Value, t.Column)
}
