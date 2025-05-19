package main

import (
    "fmt"
    "strings"
)

// GenerateFormattedTree builds a formatted string representation of the Composition
func GenerateFormattedTree(composition *Composition) string {
    var output strings.Builder
    output.WriteString("Composition\n")
    for i, paragraph := range composition.Paragraphs {
        output.WriteString(fmt.Sprintf("  Paragraph %d\n", i+1))
        
        // Upper Annotations
        output.WriteString("    Upper Annotations\n")
        for _, annotation := range paragraph.UpperAnnotations {
            output.WriteString(fmt.Sprintf("      - %s: %s\n", annotation.Type, annotation.Value))
        }

        // Letter Line
        output.WriteString("    LetterLine\n")
        if paragraph.LetterLine != nil {
            for _, element := range paragraph.LetterLine.Elements {
                if element.IsBeat {
                    output.WriteString("      - Beat:\n")
                    for _, subElement := range element.SubElements {
                        output.WriteString(fmt.Sprintf("        - %s: %s [X=%d]\n", subElement.Token.Type, subElement.Token.Value, subElement.X))
                    }
                } else {
                    output.WriteString(fmt.Sprintf("      - %s: %s [X=%d]\n", element.Token.Type, element.Token.Value, element.X))
                }
            }
        }

        // Lower Annotations
        output.WriteString("    Lower Annotations\n")
        for _, annotation := range paragraph.LowerAnnotations {
            output.WriteString(fmt.Sprintf("      - %s: %s\n", annotation.Type, annotation.Value))
        }

        // Lyrics
        output.WriteString("    Lyrics\n")
        for _, lyric := range paragraph.Lyrics {
            output.WriteString(fmt.Sprintf("      - %s: %s\n", lyric.Type, lyric.Value))
        }
    }
    return output.String()
}

