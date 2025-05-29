package parser

import "unicode"

func LexLyrics(line string) []Annotation {
	var annotations []Annotation
	start := -1
	for i, r := range line {
		if unicode.IsSpace(r) {
			if start != -1 {
				word := line[start:i]
				annotations = append(annotations, Annotation{Type: TokenTypeSyllable, Value: word, Column: start})
				start = -1
			}
		} else {
			if start == -1 {
				start = i
			}
		}
	}
	if start != -1 {
		word := line[start:]
		annotations = append(annotations, Annotation{Type: TokenTypeSyllable, Value: word, Column: start})
	}
	return annotations
}
