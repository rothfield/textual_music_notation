package main

import (
    "log"
    "strings"
    "fmt"
)

// ✅ LetterLineElement represents any top-level element (e.g., pitch, dash, barline, beat, etc.)
type LetterLineElement struct {
    Token       Token
    X           int
    SubElements []LetterLineElement // Only for Beats
    IsBeat      bool                // ✅ New flag to indicate if it's a Beat
}

// ✅ LetterLine contains all elements in a line (barlines, pitches, dashes, slurs, and beats)
type LetterLine struct {
    Elements []LetterLineElement
}


// Capitalized ParseLetterLine to be exported

func ParseLetterLine(tokens []Token) *LetterLine {

    var lineElements []LetterLineElement

    var currentBeat []LetterLineElement


    for i, token := range tokens {

        switch token.Type {

        case Dash, Pitch:

            // Collect elements as part of a beat

            currentBeat = append(currentBeat, LetterLineElement{

                Token: token,

                X:     i,

            })

        case Barline, Slur:

            // If there is an active beat, close it and add it as a LineElement

            if len(currentBeat) > 0 {

                lineElements = append(lineElements, LetterLineElement{

                    Token:       Token{Type: "Beat", Value: "Beat"},

                    X:           i,

                    SubElements: currentBeat,

                    IsBeat:      true,

                })

                currentBeat = nil

            }

            // Add the Barline or Slur to the top-level directly

            lineElements = append(lineElements, LetterLineElement{Token: token, X: i})

        default:

            // All other tokens are added directly to the top-level

            lineElements = append(lineElements, LetterLineElement{Token: token, X: i})

        }

    }


    // If there is an unfinished beat, we add it

    if len(currentBeat) > 0 {

        lineElements = append(lineElements, LetterLineElement{

            Token:       Token{Type: "Beat", Value: "Beat"},

            X:           len(tokens),

            SubElements: currentBeat,

            IsBeat:      true,

        })

    }


    return &LetterLine{

        Elements: lineElements,

    }

}

// ✅ Capitalized ParseLetterLine to be exported
func zParseLetterLine(tokens []Token) *LetterLine {
    var lineElements []LetterLineElement
    var currentBeat []LetterLineElement

    for i, token := range tokens {
        switch token.Type {
        case Dash, Pitch, Slur:
            // Collect elements as part of a beat
            currentBeat = append(currentBeat, LetterLineElement{
                Token: token,
                X:     i,
            })
        case Barline:
            // If we encounter a barline, we close the current beat (if it exists)
            if len(currentBeat) > 0 {
                // Add the collected beat as a LineElement
                lineElements = append(lineElements, LetterLineElement{
                    Token:       Token{Type: Barline, Value: "Beat"},
                    X:           i,
                    SubElements: currentBeat,
                    IsBeat:      true,
                })
                currentBeat = nil
            }
            // Add the barline separately
            lineElements = append(lineElements, LetterLineElement{Token: token, X: i})

        default:
            continue
        }
    }

    // If there is an unfinished beat, we add it
    if len(currentBeat) > 0 {
        lineElements = append(lineElements, LetterLineElement{
            Token:       Token{Type: Barline, Value: "Beat"},
            X:           len(tokens),
            SubElements: currentBeat,
            IsBeat:      true,
        })
    }

    return &LetterLine{
        Elements: lineElements,
    }
}

// ✅ ParseNotation is capitalized to be exported
func ParseNotation(notation string) string {
    log.Println("=== Begin Parsing Notation ===")
    log.Printf("Received Notation:\n%s\n", notation)

    // Lexical analysis
    tokens := Lexer(notation)
    log.Printf("Tokens generated: %v\n", tokens)

    // Parse the tokens
    letterLine := ParseLetterLine(tokens)
    log.Println("ParseLetterLine executed successfully.")

    // Build the output
    var output strings.Builder
    output.WriteString("=== Parsed Structure ===\n")

    for _, element := range letterLine.Elements {
        if element.IsBeat {
            output.WriteString("- Beat:\n")
            for _, subElement := range element.SubElements {
                output.WriteString("    - " + string(subElement.Token.Type) + ": " + subElement.Token.Value + "\n")
            }
        } else {
            output.WriteString("- " + string(element.Token.Type) + ": " + element.Token.Value + "\n")
        }
    }

    return output.String()
}

// ✅ DisplayParseTree traverses and displays the structure of LetterLine
func DisplayParseTree(letterLine *LetterLine) {
    fmt.Println("=== Parse Tree Structure ===")
    for _, element := range letterLine.Elements {
        if element.IsBeat {
            fmt.Println("- Beat:")
            for _, subElement := range element.SubElements {
                fmt.Printf("    - %s: %s\n", subElement.Token.Type, subElement.Token.Value)
            }
        } else {
            fmt.Printf("- %s: %s\n", element.Token.Type, element.Token.Value)
        }
    }
}

