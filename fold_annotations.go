package main

// ✅ FoldAnnotations traverses the parse tree and attaches annotations
// (octaves, mordents, talas, lower octaves, lyrics) to the appropriate
// LetterLine elements with ±1 column tolerance.
func FoldAnnotations(paragraph *Paragraph) {
	Log("DEBUG", "Starting FoldAnnotations")
    if paragraph.LetterLine == nil {
        return
    }

    elements := paragraph.LetterLine.Elements

    // Apply Upper Annotations (Octaves, Mordents, Talas)
    for _, token := range paragraph.UpperAnnotations {
        nearestIndex := findNearestElement(elements, token.Column)
        Log("DEBUG", "Folding annotation: %s at column %d to nearest element at index %d", token.Type, token.Column, nearestIndex)
		Log("DEBUG", "Nearest element for %s found at index %d: %v", token.Type, nearestIndex, elements[nearestIndex])
		if nearestIndex == -1 {
		Log("ERROR", "No nearest element found for annotation: %s at column %d", token.Type, token.Column)
		continue
	}
	if nearestIndex >= 0 {
            switch token.Type {
            case HigherOctave:
                elements[nearestIndex].Octave = 1
		Log("DEBUG", "Octave set to 1 for element at index %d: %v", nearestIndex, elements[nearestIndex])
            case HighestOctave:
                elements[nearestIndex].Octave = 2
		Log("DEBUG", "Octave set to 2 for element at index %d: %v", nearestIndex, elements[nearestIndex])
            case Mordent:
                elements[nearestIndex].Mordent = true
		Log("DEBUG", "Mordent set for element at index %d: %v", nearestIndex, elements[nearestIndex])
            case Tala:
                elements[nearestIndex].TalaMarker = token.Value
		Log("DEBUG", "TalaMarker set to %s for element at index %d: %v", token.Value, nearestIndex, elements[nearestIndex])
            }
        }
    }

    // Apply Lower Annotations (Lyrics and Lower Octaves)
    lyricIndex := 0
    for _, token := range paragraph.LowerAnnotations {
        nearestIndex := findNearestElement(elements, token.Column)
        Log("DEBUG", "Folding annotation: %s at column %d to nearest element at index %d", token.Type, token.Column, nearestIndex)
		Log("DEBUG", "Nearest element for %s found at index %d: %v", token.Type, nearestIndex, elements[nearestIndex])
		if nearestIndex == -1 {
		Log("ERROR", "No nearest element found for annotation: %s at column %d", token.Type, token.Column)
		continue
	}
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


