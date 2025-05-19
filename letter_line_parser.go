
package main

type LetterLineElement struct {
    Token       Token
    Column      int
    SubElements []LetterLineElement
    IsBeat      bool
    Octave      int
    Mordent     bool
    TalaMarker  string
    LyricText   string
}

// ✅ LetterLine contains all elements in a line (barlines, pitches, dashes, slurs, and beats)
type LetterLine struct {
    Elements []LetterLineElement
}

func ParseLetterLine(tokens []Token) *LetterLine {
    Log("DEBUG", "ParseLetterLine")
    Log("DEBUG", "tokens=%s", tokens)
    var lineElements []LetterLineElement
    var currentBeat []LetterLineElement

    currentColumn := 0
    for _, token := range tokens {
        switch token.Type {
        case Pitch, Dash:
            // Part of a beat, collect in currentBeat
            Log("DEBUG", "Adding to current beat: %s at column %d", token.Value, currentColumn)
            currentBeat = append(currentBeat, LetterLineElement{
                Token: token,
                Column: currentColumn,
            })
            currentColumn += len(token.Value)
        case Barline:
            // Close the current beat and add its sub-elements individually
            if len(currentBeat) > 0 {
                Log("DEBUG", "Finalizing beat with %d elements", len(currentBeat))
                for _, element := range currentBeat {
                    lineElements = append(lineElements, element)
                    Log("DEBUG", "Added element to top-level: %s at column %d", element.Token.Value, element.Column)
                }
                currentBeat = nil
            }
            // Add the barline as its own top-level element
            lineElements = append(lineElements, LetterLineElement{
                Token: token,
                Column: currentColumn,
            })
            currentColumn += len(token.Value)
        default:
            Log("WARN", "Unhandled token type in letter line: %s", token.Type)
        }
    }

    // If there is a remaining beat, unfold it into the top level
    if len(currentBeat) > 0 {
        Log("DEBUG", "Finalizing last beat with %d elements", len(currentBeat))
        for _, element := range currentBeat {
            lineElements = append(lineElements, element)
            Log("DEBUG", "Added element to top-level: %s at column %d", element.Token.Value, element.Column)
        }
    }

    return &LetterLine{
        Elements: lineElements,
    }
}
