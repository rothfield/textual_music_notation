package main

type LetterLineElement struct {
    Token       Token
    Column           int
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

    for i, token := range tokens {
        switch token.Type {
        case Pitch, Dash:
            // Start or continue a Beat
            currentBeat = append(currentBeat, LetterLineElement{
                Token: token,
                Column:     i,
            })
        case LeftSlur, RightSlur, Breath:
            if len(currentBeat) > 0 {
                // These are part of the Beat if it's active
                currentBeat = append(currentBeat, LetterLineElement{
                    Token: token,
                    Column:     i,
                })
            } else {
                // If no active Beat, they are standalone
                lineElements = append(lineElements, LetterLineElement{
                    Token: token,
                    Column:     i,
                })
            }
        case Barline:
            // Close the current Beat, if any
            if len(currentBeat) > 0 {
                lineElements = append(lineElements, LetterLineElement{
                    Token:       Token{Type: "Beat", Value: "Beat"},
                    Column:           i,
                    SubElements: currentBeat,
                    IsBeat:      true,
                })
                currentBeat = nil
            }
            // Append the Barline separately
            lineElements = append(lineElements, LetterLineElement{
                Token: token,
                Column:     i,
            })
        default:
            // Any other token type closes the Beat if active
            if len(currentBeat) > 0 {
                lineElements = append(lineElements, LetterLineElement{
                    Token:       Token{Type: "Beat", Value: "Beat"},
                    Column:           i,
                    SubElements: currentBeat,
                    IsBeat:      true,
                })
                currentBeat = nil
            }
            lineElements = append(lineElements, LetterLineElement{
                Token: token,
                Column:     i,
            })
        }
    }

    // Final check: if there's a hanging Beat, add it
    if len(currentBeat) > 0 {
        lineElements = append(lineElements, LetterLineElement{
            Token:       Token{Type: "Beat", Value: "Beat"},
            Column:           len(tokens),
            SubElements: currentBeat,
            IsBeat:      true,
        })
    }

    return &LetterLine{
        Elements: lineElements,
    }
}






