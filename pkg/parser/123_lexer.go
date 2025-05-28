package parser

func Lex123(line string) []Token {
	var tokens []Token
	col := 0
	for col < len(line) {
		ch := line[col]
		if ch == ' ' {
			col++
			continue
		}

		maxlen := 3
		if col+maxlen > len(line) {
			maxlen = len(line) - col
		}

		found := false
		for l := maxlen; l > 0; l-- {
			candidate := line[col : col+l]
			if _, ok := map[string]struct{}{
				"1":  struct{}{},
				"2":  struct{}{},
				"2#": struct{}{},
				"2b": struct{}{},
				"2x": struct{}{},
				"3":  struct{}{},
				"3#": struct{}{},
				"3b": struct{}{},
				"3x": struct{}{},
				"4":  struct{}{},
				"4#": struct{}{},
				"4b": struct{}{},
				"4x": struct{}{},
				"5":  struct{}{},
				"5#": struct{}{},
				"5b": struct{}{},
				"5x": struct{}{},
				"6":  struct{}{},
				"6#": struct{}{},
				"6b": struct{}{},
				"6x": struct{}{},
				"7":  struct{}{},
				"7#": struct{}{},
				"7b": struct{}{},
				"7x": struct{}{},
			}[candidate]; ok {
				tokens = append(tokens, Token{Type: Pitch, Value: candidate, Column: col})
				col += l
				found = true
				break
			}
		}

		if !found {
			col++
		}
	}
	return tokens
}
