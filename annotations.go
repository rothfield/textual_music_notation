package main

// UpperAnnotationLine represents the upper line of annotations (octaves, mordents, talas)
type UpperAnnotationLine struct {
    Elements []AnnotationElement
}

// LowerAnnotationLine represents the lower line of annotations (lyrics)
type LowerAnnotationLine struct {
    Elements []AnnotationElement
}

// AnnotationElement represents a single annotation (octave, mordent, tala, or lyric)
type AnnotationElement struct {
    Token Token
    X     int // X represents the column position in the line
}

// ParseUpperAnnotationLine processes the tokens for the upper annotation line
func ParseUpperAnnotationLine(tokens []Token) *UpperAnnotationLine {
    var elements []AnnotationElement
    for i, token := range tokens {
        // Only Octave, Mordent, and Tala are valid in the upper annotation line
        if token.Type == Octave || token.Type == Mordent || token.Type == Tala {
            elements = append(elements, AnnotationElement{Token: token, X: i})
        }
    }
    return &UpperAnnotationLine{
        Elements: elements,
    }
}

// ParseLowerAnnotationLine processes the tokens for the lower annotation line
func ParseLowerAnnotationLine(tokens []Token) *LowerAnnotationLine {
    var elements []AnnotationElement
    for i, token := range tokens {
        // Only Lyrics are valid in the lower annotation line
        if token.Type == Lyrics {
            elements = append(elements, AnnotationElement{Token: token, X: i})
        }
    }
    return &LowerAnnotationLine{
        Elements: elements,
    }
}

// Display method for debugging and visualization
func (ual *UpperAnnotationLine) Display() {
    println("=== Upper Annotation Line ===")
    for _, element := range ual.Elements {
        println("Element:", element.Token.Value, "at position", element.X)
    }
}

func (lal *LowerAnnotationLine) Display() {
    println("=== Lower Annotation Line ===")
    for _, element := range lal.Elements {
        println("Element:", element.Token.Value, "at position", element.X)
    }
}

