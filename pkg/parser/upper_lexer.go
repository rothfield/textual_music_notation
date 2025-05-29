package parser

func LexUpper(line string) []Annotation {
	var annotations []Annotation
	Log("DEBUG", "Entering LexUpper:")
	column := 0
	for _, r := range line {
		Log("DEBUG", "LexUpper: rune=%q col=%d", r, column)
		switch r {
		case '.':
			annotations = append(annotations, Annotation{Type: TokenTypeUpperOctave, Value: ".", Column: column})
		case '~':
			annotations = append(annotations, Annotation{Type: TokenTypeMordent, Value: "~", Column: column})
		case ':':
			annotations = append(annotations, Annotation{Type: TokenTypeHighestOctave, Value: ":", Column: column})
		case '0', '1', '2', '3', '4', '5', '6', '7', '8':
			annotations = append(annotations, Annotation{Type: TokenTypeTala, Value: string(r), Column: column})
		}
		column++
	}
	return annotations
}
