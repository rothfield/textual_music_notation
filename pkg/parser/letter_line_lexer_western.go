package parser

func LexLetterLineWestern(line string) []Token {
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
				"A":  struct{}{},
				"A#": struct{}{},
				"Ab": struct{}{},
				"Ax": struct{}{},
				"B":  struct{}{},
				"B#": struct{}{},
				"Bb": struct{}{},
				"Bx": struct{}{},
				"C":  struct{}{},
				"C#": struct{}{},
				"Cb": struct{}{},
				"Cx": struct{}{},
				"D":  struct{}{},
				"D#": struct{}{},
				"Db": struct{}{},
				"Dx": struct{}{},
				"E":  struct{}{},
				"E#": struct{}{},
				"Eb": struct{}{},
				"Ex": struct{}{},
				"F":  struct{}{},
				"F#": struct{}{},
				"Fb": struct{}{},
				"Fx": struct{}{},
				"G":  struct{}{},
				"G#": struct{}{},
				"Gb": struct{}{},
				"Gx": struct{}{},
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
