package main

// ✅ FoldAnnotations traverses the parse tree and attaches annotations
// (octaves, mordents, talas, lower octaves, lyrics) to the appropriate
// LetterLine elements with ±1 column tolerance.
func FoldAnnotations(paragraph *Paragraph) {
    if paragraph.LetterLine == nil {
        return
    }

    elements := paragraph.LetterLine.Elements

    // Apply Upper Annotations (Octaves, Mordents, Talas)
    for _, token := range paragraph.UpperAnnotations {
        nearestIndex := findNearestElement(elements, token.Column)
        if nearestIndex >= 0 {
            switch token.Type {
            case HigherOctave:
                elements[nearestIndex].Octave = 1
            case HighestOctave:
                elements[nearestIndex].Octave = 2
            case Mordent:
                elements[nearestIndex].Mordent = true
            case Tala:
                elements[nearestIndex].TalaMarker = token.Value
            }
        }
    }

    // Apply Lower Annotations (Lyrics and Lower Octaves)
    lyricIndex := 0
    for _, token := range paragraph.LowerAnnotations {
        nearestIndex := findNearestElement(elements, token.Column)
        if nearestIndex >= 0 {
            if token.Type == LowerOctave {
                elements[nearestIndex].Octave = -1
            }
        }
    }

    // Apply Lyrics to corresponding elements
    for lyricIndex < len(paragraph.Lyrics) && lyricIndex < len(elements) {
        elements[lyricIndex].LyricText = paragraph.Lyrics[lyricIndex].Value
        lyricIndex++
    }

    // Overflow lyrics map to the last pitch
    for lyricIndex < len(paragraph.Lyrics) {
        elements[len(elements)-1].LyricText += " " + paragraph.Lyrics[lyricIndex].Value
        lyricIndex++
    }
}

// findNearestElement finds the nearest element in the LetterLine
// within ±1 column of the given position.
func findNearestElement(elements []LetterLineElement, column int) int {
    bestIndex := -1
    bestDistance := 2 // We only care about ±1

    for i, elem := range elements {
        dist := abs(elem.Column - column)
        if dist < bestDistance {
            bestDistance = dist
            bestIndex = i
        }
    }
    return bestIndex
}


