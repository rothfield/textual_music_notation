package main

import "strings"

// ✅ Composition represents the entire parsed structure
type Composition struct {
    Paragraphs []Paragraph
}

// ✅ ParseComposition parses the input notation into a Composition structure
func ParseComposition(input string) *Composition {
    Log("DEBUG", "ParseComposition")
    
    // Split the input into paragraphs by double newline
    paragraphs := strings.Split(input, "\n\n")
    var parsedParagraphs []Paragraph

    // Parse each paragraph separately
    for _, para := range paragraphs {
        paragraph := ParseParagraph(para)
        parsedParagraphs = append(parsedParagraphs, paragraph)
    }

    // Return the full composition structure
    return &Composition{
        Paragraphs: parsedParagraphs,
    }
}

