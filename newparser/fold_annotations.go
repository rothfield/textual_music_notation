package newparser

func FoldAnnotations(p *Paragraph, annotations []Annotation) {
    if p == nil || p.LetterLine == nil {
        return
    }

    for _, ann := range annotations {
        best := -1
        bestDist := 2 // only match within ±1 column

        for i, el := range p.LetterLine.Elements {
            dist := abs(el.Column - ann.Column)
            if dist < bestDist {
                best = i
                bestDist = dist
            }
        }

        if best == -1 {
            continue
        }

        el := &p.LetterLine.Elements[best]

        switch ann.Type {
        case UpperOctave:
            el.Octave += 1
        case LowerOctave:
            el.Octave -= 1
        case Mordent:
            el.Mordent = true
        case Syllable:
            el.SyllableText = ann.Value
        case Tala:
            el.TalaMarker = ann.Value
        }
    }
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
