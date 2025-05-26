package parser

func LexLetterLineSargam(input string) []Token {
	var tokens []Token
	i := 0
	for i < len(input) {

		// Check for 4-character barline
		if i+3 < len(input) && input[i:i+4] == ":||:" {
			tokens = append(tokens, Token{Type: Barline, Value: ":||:", Column: i})
			i += 4
			continue
		}

		// Check for 2-character barlines
		if i+1 < len(input) {
			pair := input[i : i+2]
			switch pair {
			case ".|":
				tokens = append(tokens, Token{Type: Barline, Value: ".|", Column: i})
				i += 2
				continue
			case ":|":
				tokens = append(tokens, Token{Type: Barline, Value: ":|", Column: i})
				i += 2
				continue
			case ":|:":
				tokens = append(tokens, Token{Type: Barline, Value: ":|:", Column: i})
				i += 2
				continue
			case ":||:":
				tokens = append(tokens, Token{Type: Barline, Value: ":||:", Column: i})
				i += 2
				continue
			case "[|":
				tokens = append(tokens, Token{Type: Barline, Value: "[|", Column: i})
				i += 2
				continue
			case "|.":
				tokens = append(tokens, Token{Type: Barline, Value: "|.", Column: i})
				i += 2
				continue
			case "|:":
				tokens = append(tokens, Token{Type: Barline, Value: "|:", Column: i})
				i += 2
				continue
			case "|]":
				tokens = append(tokens, Token{Type: Barline, Value: "|]", Column: i})
				i += 2
				continue
			case "||":
				tokens = append(tokens, Token{Type: Barline, Value: "||", Column: i})
				i += 2
				continue
			}
		}

		// Check for multi-char pitches
		if i+1 < len(input) && input[i:i+2] == "P#" {
			tokens = append(tokens, Token{Type: Pitch, Value: "P#", Column: i})
			i += 2
			continue
		}
		if i+1 < len(input) && input[i:i+2] == "D#" {
			tokens = append(tokens, Token{Type: Pitch, Value: "D#", Column: i})
			i += 2
			continue
		}
		if i+1 < len(input) && input[i:i+2] == "R#" {
			tokens = append(tokens, Token{Type: Pitch, Value: "R#", Column: i})
			i += 2
			continue
		}
		if i+1 < len(input) && input[i:i+2] == "S#" {
			tokens = append(tokens, Token{Type: Pitch, Value: "S#", Column: i})
			i += 2
			continue
		}
		if i+1 < len(input) && input[i:i+2] == "Gb" {
			tokens = append(tokens, Token{Type: Pitch, Value: "Gb", Column: i})
			i += 2
			continue
		}
		char := input[i]
		switch char {
		case 'S', 'r', 'R', 'g', 'G', 'm', 'M', 'd', 'D', 'n', 'N':
			tokens = append(tokens, Token{Type: Pitch, Value: string(char), Column: i})
		case '-':
			tokens = append(tokens, Token{Type: Dash, Value: string(char), Column: i})
		case '|':
			tokens = append(tokens, Token{Type: Barline, Value: string(char), Column: i})
		case '\'':
			tokens = append(tokens, Token{Type: Breath, Value: string(char), Column: i})
		case '(':
			tokens = append(tokens, Token{Type: LeftSlur, Value: string(char), Column: i})
		case ')':
			tokens = append(tokens, Token{Type: RightSlur, Value: string(char), Column: i})
		default:
			tokens = append(tokens, Token{Type: Unknown, Value: string(char), Column: i})
		}
		i++
	}
	return tokens
}
