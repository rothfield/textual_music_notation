package main

import (
	"log"
	"regexp"
	"strings"
)

// TokenRegex represents a mapping of regex patterns to token types
type TokenRegex struct {
	Pattern *regexp.Regexp
	Type    TokenType
}

// Lexer tokenizes the letter notation into individual tokens.
func LetterLineLexer(notation string) []Token {
	var tokens []Token
	notation = strings.TrimSpace(notation)

	// Define all token patterns in a readable structure
	tokenPatterns := []TokenRegex{
		{Pattern: regexp.MustCompile(`\|:|\|\||\|`), Type: Barline},
		{Pattern: regexp.MustCompile(`-`), Type: Dash},
		{Pattern: regexp.MustCompile(`\(`), Type: LeftSlur},
		{Pattern: regexp.MustCompile(`\)`), Type: RightSlur},
		{Pattern: regexp.MustCompile(`'`), Type: Breath},
		{Pattern: regexp.MustCompile(`[:;]`), Type: Octave},
	//	{Pattern: regexp.MustCompile(`~`), Type: Mordent},
	//	{Pattern: regexp.MustCompile(`[0-8+.]`), Type: Tala},
	//	{Pattern: regexp.MustCompile(`[a-zA-Z]+`), Type: Lyrics},
		{Pattern: regexp.MustCompile(`\s+`), Type: Space},
		{Pattern: regexp.MustCompile(`[SrRgGmMPdDnN]`), Type: Pitch},
	}

	// Lexical Analysis Loop
	for len(notation) > 0 {
		matched := false
		for _, tokenRegex := range tokenPatterns {
			if loc := tokenRegex.Pattern.FindStringIndex(notation); loc != nil && loc[0] == 0 {
				match := notation[loc[0]:loc[1]]
				tokens = append(tokens, Token{Type: tokenRegex.Type, Value: match})
				notation = notation[len(match):]
				matched = true
				break
			}
		}
		if !matched {
			log.Printf("Unrecognized token at: %s\n", notation)
			notation = notation[1:]
		}
	}

	return tokens
}

