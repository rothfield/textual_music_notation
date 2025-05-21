package newparser

import "unicode"

func LexUpperAnnotationLine(line string) []Annotation {
    var annotations []Annotation
    for i, r := range line {
        switch r {
        case '~':
            annotations = append(annotations, Annotation{Type: Mordent, Value: "~", Column: i})
        case ':':
            annotations = append(annotations, Annotation{Type: UpperOctave, Value: ":", Column: i})
        case '0', '1', '2', '3', '4', '5', '6', '7', '8':
            annotations = append(annotations, Annotation{Type: Tala, Value: string(r), Column: i})
        }
    }
    return annotations
}

func LexLowerAnnotationLine(line string) []Annotation {
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

func LexSyllableAnnotationLine(line string) []Annotation {
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
    return annotations
}
