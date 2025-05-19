package main

import (
    "fmt"
)

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
                    formatter.WriteSubElement(indent, string(subElement.Token.Type), subElement.Token.Value, subElement.X)  // Convert TokenType to string
                }
            } else {
                formatter.WriteElement(indent, string(element.Token.Type), element.Token.Value, element.X)  // Convert TokenType to string
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

