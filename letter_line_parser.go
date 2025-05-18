package main

import (
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
                X:     i,
            })
        case LeftSlur, RightSlur, Breath:
            if len(currentBeat) > 0 {
                // These are part of the Beat if it's active
                currentBeat = append(currentBeat, LetterLineElement{
                    Token: token,
                    X:     i,
                })
            } else {
                // If no active Beat, they are standalone
                lineElements = append(lineElements, LetterLineElement{
                    Token: token,
                    X:     i,
                })
            }
        case Barline:
            // Close the current Beat, if any
            if len(currentBeat) > 0 {
                lineElements = append(lineElements, LetterLineElement{
                    Token:       Token{Type: "Beat", Value: "Beat"},
                    X:           i,
                    SubElements: currentBeat,
                    IsBeat:      true,
                })
                currentBeat = nil
            }
            // Append the Barline separately
            lineElements = append(lineElements, LetterLineElement{
                Token: token,
                X:     i,
            })
        default:
            // Any other token type closes the Beat if active
            if len(currentBeat) > 0 {
                lineElements = append(lineElements, LetterLineElement{
                    Token:       Token{Type: "Beat", Value: "Beat"},
                    X:           i,
                    SubElements: currentBeat,
                    IsBeat:      true,
                })
                currentBeat = nil
            }
            lineElements = append(lineElements, LetterLineElement{
                Token: token,
                X:     i,
            })
        }
    }

    // Final check: if there's a hanging Beat, add it
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





// ✅ ParseNotation is capitalized to be exported
func ParseNotation(notation string) string {
    Log("DEBUG","=== Begin Parsing Notation ===")
   Log("DEBUG","Received Notation:\n%s\n", notation)

    // Lexical analysis
    tokens := Lexer(notation)
    Log("DEBUG","Tokens generated: %v\n", tokens)

    // Parse the tokens
    letterLine := ParseLetterLine(tokens)
    Log("DEBUG","ParseLetterLine executed successfully.")

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



// Test cases to validate the parser
func TestParseLetterLine() {
    testCases := []string{
        "|S- -- -(-  S-r )",
    }

    for _, testCase := range testCases {
        fmt.Println("--- Testing: ", testCase)
        tokens := Lexer(testCase)
        letterLine := ParseLetterLine(tokens)
        DisplayParseTree(letterLine)
    }
}

func zmain() {
	InitLogger()
	defer logFile.Close()
	Log("DEBUG","Test log message from main")
  fmt.Println("If you see this, logger should be working.")
  TestParseLetterLine()
}


