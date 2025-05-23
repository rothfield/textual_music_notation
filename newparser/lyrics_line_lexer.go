package newparser

import "unicode"

func LexLyricsAnnotationLine(line string) []Annotation {
    var annotations []Annotation
    start := -1
    for i, r := range line {
        if unicode.IsSpace(r) {
            if start != -1 {
                word := line[start:i]
                annotations = append(annotations, Annotation{Type: Syllable, Value: word, Column: start})
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
        annotations = append(annotations, Annotation{Type: Syllable, Value: word, Column: start})
    }
		Log("DEBUG", "Lexed %d syllables", len(annotations))
    return annotations
}

