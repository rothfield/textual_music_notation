package main

import (
    "fmt"
    "strings"
)

// GenerateFormattedTree generates the formatted tree for the entire composition
func GenerateFormattedTree(composition *Composition) string {
    formatter := &StringFormatter{}

    // ✅ Display the raw text directly, left-justified with "Raw Text:" header
    if composition.RawText != "" {
        formatter.WriteLine("  ", "Raw Text:")
        rawLines := strings.Split(composition.RawText, "\n")
        for _, line := range rawLines {
            formatter.WriteLine("  ", line)  // Left-justified but indented under "Raw Text:"
        }
    }

    // ✅ Display the parsed tree as usual
    formatter.WriteLine("", "Composition")
    for i, paragraph := range composition.Paragraphs {
        formatter.WriteLine("  ", fmt.Sprintf("Paragraph %d", i+1))
        RenderParagraph(paragraph, formatter, "    ")
    }
    return formatter.Builder.String()
}

// GenerateFormattedTree generates the formatted tree for the entire composition
func xzzGenerateFormattedTree(composition *Composition) string {
	
    formatter := &StringFormatter{}

    // ✅ Display the raw text directly, left-justified with no indentation
    if composition.RawText != "" {
        formatter.WriteLine("", "=== Raw Notation ===")
        rawLines := strings.Split(composition.RawText, "\n")
        for _, line := range rawLines {
            formatter.WriteLine("", line)  // No indent, left-justified
        }
        formatter.WriteLine("", "====================")
    }

    // ✅ Display the parsed tree as usual
    formatter.WriteLine("", "Composition")
    for i, paragraph := range composition.Paragraphs {
        formatter.WriteLine("  ", fmt.Sprintf("Paragraph %d", i+1))
        RenderParagraph(paragraph, formatter, "    ")
    }
    return formatter.Builder.String()
}


// GenerateFormattedTree generates the formatted tree for the entire composition
func zGenerateFormattedTree(composition *Composition) string {
    formatter := &StringFormatter{}

    // Display the raw text directly, no indentation
    if composition.RawText != "" {
        formatter.WriteLine("", "=== Raw Notation ===")
        formatter.WriteLine("", composition.RawText)
        formatter.WriteLine("", "====================")
    }

    // Display the parsed tree as usual
    formatter.WriteLine("", "Composition")
    for i, paragraph := range composition.Paragraphs {
        formatter.WriteLine("  ", fmt.Sprintf("Paragraph %d", i+1))
        RenderParagraph(paragraph, formatter, "    ")
    }
    return formatter.Builder.String()
}

// RenderParagraph renders the details of a Paragraph using the given formatter
func RenderParagraph(paragraph Paragraph, formatter *StringFormatter, indent string) {
    // Upper Annotations (only if not nil and not empty)
    if len(paragraph.UpperAnnotations) > 0 {
        formatter.WriteLine(indent, "Upper Annotations")
        for _, annotation := range paragraph.UpperAnnotations {
            formatter.WriteAnnotation(indent, string(annotation.Type), annotation.Value, annotation.Column)  
        }
    }

    // LetterLine (always displayed)
    formatter.WriteLine(indent, "LetterLine")
    if paragraph.LetterLine != nil {
        for _, element := range paragraph.LetterLine.Elements {
            if element.IsBeat {
                formatter.WriteLine(indent, "  - Beat:")
                for _, subElement := range element.SubElements {
                    formatter.WriteSubElement(indent, string(subElement.Token.Type), subElement.Token.Value, subElement.Column, subElement.Octave, subElement.Mordent, subElement.TalaMarker, subElement.LyricText)  
                }
            } else {
                formatter.WriteElement(indent, string(element.Token.Type), element.Token.Value, element.Column, element.Octave, element.Mordent, element.TalaMarker, element.LyricText)  
            }
        }
    }

    // Lower Annotations (only if not nil and not empty)
    if len(paragraph.LowerAnnotations) > 0 {
        formatter.WriteLine(indent, "Lower Annotations")
        for _, annotation := range paragraph.LowerAnnotations {
            formatter.WriteAnnotation(indent, string(annotation.Type), annotation.Value, annotation.Column)  
        }
    }

    // Lyrics (only if not nil and not empty)
    if len(paragraph.Lyrics) > 0 {
        formatter.WriteLine(indent, "Lyrics")
        for _, lyric := range paragraph.Lyrics {
            formatter.WriteAnnotation(indent, string(lyric.Type), lyric.Value, lyric.Column)  
        }
    }
}

