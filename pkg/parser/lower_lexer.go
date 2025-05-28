package parser

func LexLower(line string) []Annotation {
	var annotations []Annotation
	for i, r := range line {
		switch r {
		case '.':
			annotations = append(annotations, Annotation{Type: LowerOctave, Value: ".", Column: i})
		case '0', '1', '2', '3', '4', '5', '6', '7', '8':
			annotations = append(annotations, Annotation{Type: Tala, Value: string(r), Column: i})
		}
	}
	return annotations
}
