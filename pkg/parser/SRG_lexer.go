package parser

import (
	"regexp"
	"textual_music_notation/internal/logger"
)

var whitespaceRE = regexp.MustCompile(`^[ \t]+`)

func LexLineSargam(input string) []Token {
	logger.Log("DEBUG", "Entering LexLineSargm, string is %s", input)
	var tokens []Token
	i := 0
	for i < len(input) {
		remaining := input[i:]

		if m := whitespaceRE.FindString(remaining); m != "" {
			tokens = append(tokens, Token{Type: TokenTypeSpace, Value: m, Column: i})
			i += len(m)
			continue
		}

		// Check for 4-character barline
		if i+3 < len(input) && input[i:i+4] == ":||:" {
			tokens = append(tokens, Token{Type: TokenTypeBarline, Value: ":||:", Column: i})
			i += 4
			continue
		}

		// Check for 2-character barlines
		if i+1 < len(input) {
			pair := input[i : i+2]
			switch pair {
			case ".|", ":|", ":|:", ":||:", "[|", "|.", "|:", "|]", "||":
				tokens = append(tokens, Token{Type: TokenTypeBarline, Value: pair, Column: i})
				i += 2
				continue
			}
		}

		// Check for multi-char pitches
		if i+1 < len(input) {
			two := input[i : i+2]
			switch two {
			case "P#", "D#", "R#", "S#", "Gb":
				tokens = append(tokens, Token{Type: TokenTypePitch, Value: two, Column: i})
				i += 2
				continue
			}
		}

		char := input[i]
		switch char {
		case 'S', 'r', 'R', 'g', 'G', 'm', 'M', 'd', 'D', 'n', 'N':
			tokens = append(tokens, Token{Type: TokenTypePitch, Value: string(char), Column: i})
		case '-':
			tokens = append(tokens, Token{Type: TokenTypeDash, Value: string(char), Column: i})
		case '|':
			tokens = append(tokens, Token{Type: TokenTypeBarline, Value: string(char), Column: i})
		case '\'':
			tokens = append(tokens, Token{Type: TokenTypeBreath, Value: string(char), Column: i})
		case '(':
			tokens = append(tokens, Token{Type: TokenTypeLeftSlur, Value: string(char), Column: i})
		case ')':
			tokens = append(tokens, Token{Type: TokenTypeRightSlur, Value: string(char), Column: i})
		default:
			tokens = append(tokens, Token{Type: TokenTypeUnknown, Value: string(char), Column: i})
		}
		i++
	}
	logger.Log("ERROR", "*****leaving lexer, tokens= %s", tokens)
	return tokens
}
