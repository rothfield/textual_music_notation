package newparser

func FoldAnnotations(p *Paragraph, annotations []Annotation) {
	if p == nil || p.LetterLine == nil {
		Log("DEBUG", "FoldAnnotations skipped: nil paragraph or letter line")
		return
	}

	Log("DEBUG", "FoldAnnotations called with %d annotations", len(annotations))

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
			Log("DEBUG", "No match found for annotation %s at column %d", ann.Type, ann.Column)
			continue
		}

		el := &p.LetterLine.Elements[best]

		// If it's a beat, try folding to its inner elements
		if el.IsBeat && len(el.SubElements) > 0 {
			bestInner := -1
			bestInnerDist := 2
			for i, sub := range el.SubElements {
				dist := abs(sub.Column - ann.Column)
				if dist < bestInnerDist {
					bestInner = i
					bestInnerDist = dist
				}
			}
			if bestInner != -1 {
				Log("DEBUG", "Folding annotation %s at column %d to inner beat element at index %d (column %d)",
					ann.Type, ann.Column, bestInner, el.SubElements[bestInner].Column)
				applyAnnotation(&el.SubElements[bestInner], ann)
				continue
			}
		}

		Log("DEBUG", "Folding annotation %s at column %d to element at index %d (column %d)",
			ann.Type, ann.Column, best, el.Column)
		Log("DEBUG", "Annotation %s at col %d not applied — no subelement match in beat", ann.Type, ann.Column)
	}
}

func applyAnnotation(el *LetterLineElement, ann Annotation) {
		Log("DEBUG", "applyAnnotation, annotation= %s ",ann)
	switch ann.Type {
	case HigherOctave:
		el.Octave += 1
	case LowerOctave:
		el.Octave -= 1
	case Mordent:
		el.Mordent = true
	case Syllable:
		el.Syllables = append(el.Syllables, ann.Value)
	case TalaMarker:
		el.TalaMarker = ann.Value
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

