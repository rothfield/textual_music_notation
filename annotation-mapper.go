package main

// AnnotationMapper maps annotations to corresponding line elements
type AnnotationMapper struct {
    UpperAnnotations []AnnotationElement
    LowerAnnotations []AnnotationElement
}

// ✅ MapAnnotations applies annotations (octave, mordent, lyric) to the letter line
func (am *AnnotationMapper) MapAnnotations(letterLine *LetterLine) {
    // Map Upper Annotations (Octave, Mordent, Tala)
    for _, annotation := range am.UpperAnnotations {
        for i, element := range letterLine.Elements {
            if element.IsBeat {
                // Traverse SubElements inside the Beat
                for j, subElement := range element.SubElements {
                    if subElement.Token.Type == Pitch && abs(subElement.X-annotation.X) <= 1 {
                        switch annotation.Token.Type {
                        case Octave:
                            subElement.Token.Value = annotation.Token.Value + subElement.Token.Value
                        case Mordent:
                            subElement.Token.Value += "~"
                        case Tala:
                            // Talas do not affect the pitch itself, but could be marked as metadata
                        }
                        letterLine.Elements[i].SubElements[j] = subElement
                    }
                }
            } else {
                // If it's a top-level element, map directly
                if element.Token.Type == Pitch && abs(element.X-annotation.X) <= 1 {
                    switch annotation.Token.Type {
                    case Octave:
                        element.Token.Value = annotation.Token.Value + element.Token.Value
                    case Mordent:
                        element.Token.Value += "~"
                    case Tala:
                        // Talas can be applied here as well
                    }
                    letterLine.Elements[i] = element
                }
            }
        }
    }

    // Map Lower Annotations (Lyrics)
    for _, annotation := range am.LowerAnnotations {
        for i, element := range letterLine.Elements {
            if element.IsBeat {
                for j, subElement := range element.SubElements {
                    if subElement.Token.Type == Pitch && abs(subElement.X-annotation.X) <= 1 {
                        subElement.Token.Value += "[" + annotation.Token.Value + "]"
                        letterLine.Elements[i].SubElements[j] = subElement
                    }
                }
            } else {
                if element.Token.Type == Pitch && abs(element.X-annotation.X) <= 1 {
                    element.Token.Value += "[" + annotation.Token.Value + "]"
                    letterLine.Elements[i] = element
                }
            }
        }
    }
}

