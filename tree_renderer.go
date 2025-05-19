package main

import (
    "fmt"
)
// ✅ DisplayParseTree traverses and displays the structure of LetterLine
func DisplayParseTree(letterLine *LetterLine) {
    fmt.Println("=== Parse Tree Structure ===")
    for _, element := range letterLine.Elements {
        if element.IsBeat {
            fmt.Println("- Beat:")
            for _, subElement := range element.SubElements {
                fmt.Printf("    - %s: %s [Column=%d]\n", subElement.Token.Type, subElement.Token.Value, subElement.Column)
                if subElement.Octave != 0 {
                    fmt.Printf("      - Octave: %d\n", subElement.Octave)
                }
                if subElement.Mordent {
                    fmt.Println("      - Mordent: true")
                }
                if subElement.TalaMarker != "" {
                    fmt.Printf("      - Tala: %s\n", subElement.TalaMarker)
                }
                if subElement.LyricText != "" {
                    fmt.Printf("      - Lyric: %s\n", subElement.LyricText)
                }
            }
        } else {
            fmt.Printf("- %s: %s [Column=%d]\n", element.Token.Type, element.Token.Value, element.Column)
            if element.Octave != 0 {
                fmt.Printf("  - Octave: %d\n", element.Octave)
            }
            if element.Mordent {
                fmt.Println("  - Mordent: true")
            }
            if element.TalaMarker != "" {
                fmt.Printf("  - Tala: %s\n", element.TalaMarker)
            }
            if element.LyricText != "" {
                fmt.Printf("  - Lyric: %s\n", element.LyricText)
            }
        }
    }
}

// RenderParagraph renders the details of a Paragraph using the given formatter
func RenderParagraph(paragraph Paragraph, formatter *StringFormatter, indent string) {
    // Upper Annotations (only if not nil and not empty)
    if len(paragraph.UpperAnnotations) > 0 {
        formatter.WriteLine(indent, "Upper Annotations")
        for _, annotation := range paragraph.UpperAnnotations {
            formatter.WriteAnnotation(indent, string(annotation.Type), annotation.Value)  // Convert TokenType to string
        }
    }

    // LetterLine (always displayed)
    formatter.WriteLine(indent, "LetterLine")
    if paragraph.LetterLine != nil {
        for _, element := range paragraph.LetterLine.Elements {
            if element.IsBeat {
                formatter.WriteLine(indent, "  - Beat:")
                for _, subElement := range element.SubElements {
                    formatter.WriteSubElement(indent, string(subElement.Token.Type), subElement.Token.Value, subElement.Column)  // Convert TokenType to string
                }
            } else {
                formatter.WriteElement(indent, string(element.Token.Type), element.Token.Value, element.Column)  // Convert TokenType to string
            }
        }
    }

    // Lower Annotations (only if not nil and not empty)
    if len(paragraph.LowerAnnotations) > 0 {
        formatter.WriteLine(indent, "Lower Annotations")
        for _, annotation := range paragraph.LowerAnnotations {
            formatter.WriteAnnotation(indent, string(annotation.Type), annotation.Value)  // Convert TokenType to string
        }
    }

    // Lyrics (only if not nil and not empty)
    if len(paragraph.Lyrics) > 0 {
        formatter.WriteLine(indent, "Lyrics")
        for _, lyric := range paragraph.Lyrics {
            formatter.WriteAnnotation(indent, string(lyric.Type), lyric.Value)  // Convert TokenType to string
        }
    }
}

// GenerateFormattedTree generates the formatted tree for the entire composition
func GenerateFormattedTree(composition *Composition) string {
    formatter := &StringFormatter{}
    formatter.WriteLine("", "Composition")
    for i, paragraph := range composition.Paragraphs {
        formatter.WriteLine("  ", fmt.Sprintf("Paragraph %d", i+1))
        RenderParagraph(paragraph, formatter, "    ")
    }
    return formatter.Builder.String()
}


