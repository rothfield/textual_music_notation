package newparser

func LexUpperAnnotationLine(line string) []Annotation {
    var annotations []Annotation
    Log("DEBUG", "Entering LexUpperAnnotationLine:" )
    column := 0
    for _, r := range line {
        Log("DEBUG", "LexUpperAnnotationLine: rune=%q col=%d", r, column)
        switch r {
        case '.':
            annotations = append(annotations, Annotation{Type: UpperOctave, Value: ".", Column: column})
        case '~':
            annotations = append(annotations, Annotation{Type: Mordent, Value: "~", Column: column})
        case ':':
            annotations = append(annotations, Annotation{Type: HighestOctave, Value: ":", Column: column})
        case '0', '1', '2', '3', '4', '5', '6', '7', '8':
            annotations = append(annotations, Annotation{Type: Tala, Value: string(r), Column: column})
        }
        column++
    }
    return annotations
}

