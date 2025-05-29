package parser

func Walk(line *Line, visit func(*Element)) {
	for i := range line.Elements {
		el := &line.Elements[i]
		if el.IsBeat {
			for j := range el.SubElements {
				visit(&el.SubElements[j])
			}
		} else {
			visit(el)
		}
	}
}

func ApplyAnnotations(p *Paragraph, annotations []Annotation) {
	if p == nil || p.Line == nil {
		return
	}

	Log("DEBUG", "ApplyAnnotations called with %d annotations", len(annotations))
	Log("DEBUG", "ApplyAnnotations called with annotations: %s", annotations)

	// Flatten all pitch/dash-level elements with their positions
	var elements []*Element
	Walk(p.Line, func(el *Element) {
		elements = append(elements, el)
	})

	for _, ann := range annotations {
		best := -1
		bestDist := 1000
		for i, el := range elements {
			dist := abs(el.Column - ann.Column)

			if ann.Type == TokenTypeSyllable {
				// Relaxed syllable: match nearest pitch only
				if el.Token.Type == TokenTypePitch && dist < bestDist {
					best = i
					bestDist = dist
				}
			} else {
				// Strict matching for other annotations
				if dist <= 1 && dist < bestDist {
					best = i
					bestDist = dist
				}
			}
		}

		if best != -1 {
			el := elements[best]
			Log("DEBUG", "Folding annotation %s at column %d to element at col %d", ann.Type, ann.Column, el.Column)
			applyAnnotation(el, ann)
		} else if ann.Type == TokenTypeSyllable {
			el := fallbackToLastPitch(p.Line, ann)
			if el != nil {
				Log("DEBUG", "Fallback applied for syllable at column %d", ann.Column)
				applyFallbackAnnotation(el, ann)
			}
		} else {
			Log("DEBUG", "No match found for annotation %s at column %d", ann.Type, ann.Column)
		}
	}
}

func applyAnnotation(el *Element, ann Annotation) {
	Log("DEBUG", "applyAnnotation, annotation= Type=%s, Value=%s, Column=%d", ann.Type, ann.Value, ann.Column)
	switch ann.Type {
	case TokenTypeUpperOctave:
		el.Octave = 1
	case TokenTypeHighestOctave:
		el.Octave = 2
	case TokenTypeLowerOctave:
		el.Octave = -1
	case TokenTypeLowestOctave:
		el.Octave = -2
	case TokenTypeMordent:
		el.Mordent = true
	case TokenTypeSyllable:
		if el.Syllable == "" {
			el.Syllable = ann.Value
		} else {
			el.ExtraSyllables = append(el.ExtraSyllables, ann.Value)
		}
	case TokenTypeTala:
		el.Tala = ann.Value
	}
}

func fallbackToLastPitch(line *Line, ann Annotation) *Element {
	for i := len(line.Elements) - 1; i >= 0; i-- {
		el := &line.Elements[i]
		if el.IsBeat {
			for j := len(el.SubElements) - 1; j >= 0; j-- {
				sub := &el.SubElements[j]
				if sub.Token.Type == TokenTypePitch {
					return sub
				}
			}
		} else {
			if el.Token.Type == TokenTypePitch {
				return el
			}
		}
	}
	return nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func applyFallbackAnnotation(el *Element, ann Annotation) {
	Log("DEBUG", "applyFallbackAnnotation: %s -> ExtraSyllables", ann.Value)
	el.ExtraSyllables = append(el.ExtraSyllables, ann.Value)
}

