package newparser

import (
    "fmt"
    "strings"
)

type StringFormatter struct {
    Builder strings.Builder
}

func (f *StringFormatter) WriteLine(indent, line string) {
    f.Builder.WriteString(indent + line + "\n")
}

func RenderLetterLine(line *LetterLine, formatter *StringFormatter, indent string) {
    formatter.WriteLine(indent, "LetterLine")
    for _, el := range line.Elements {
        if el.IsBeat {
            formatter.WriteLine(indent+"  ", "- Beat:")
            for _, sub := range el.SubElements {
                formatter.WriteLine(indent+"    ", "- "+describeToken(sub))
            }
        } else {
            formatter.WriteLine(indent+"  ", "- "+describeToken(el))
        }
    }
}

func describeToken(el LetterLineElement) string {
    Log("DEBUG", "describeToken called for %s at column %d (Octave=%d)", el.Token.Value, el.Column, el.Octave)

    parts := []string{fmt.Sprintf("%s: %s [Column=%d]", string(el.Token.Type), el.Token.Value, el.Column)}
    if el.Octave != 0 {
        parts = append(parts, fmt.Sprintf("Octave: %+d", el.Octave))
    }
    if el.Mordent {
        parts = append(parts, "Mordent: true")
    }
    if el.TalaMarker != "" {
        parts = append(parts, fmt.Sprintf("Tala: %s", el.TalaMarker))
    }
    if el.SyllableText != "" {
        parts = append(parts, fmt.Sprintf("Lyric: %s", el.SyllableText))
    }
    return strings.Join(parts, ", ")
}

func RenderParagraph(p *Paragraph, formatter *StringFormatter) {
    formatter.WriteLine("", "=== Raw Notation ===")
    formatter.WriteLine("", p.Raw)
    formatter.WriteLine("", "====================")

    if len(p.UpperAnnotations) > 0 {
        formatter.WriteLine("", "  Upper Annotations")
        for _, l := range p.UpperAnnotations {
            formatter.WriteLine("    ", l)
        }
    }

    RenderLetterLine(p.LetterLine, formatter, "  ")

    if len(p.LowerAnnotations) > 0 {
        formatter.WriteLine("", "  Lower Annotations")
        for _, l := range p.LowerAnnotations {
            formatter.WriteLine("    ", l)
        }
    }

    if len(p.Lyrics) > 0 {
        formatter.WriteLine("", "  Lyrics")
        for _, l := range p.Lyrics {
            formatter.WriteLine("    ", l)
        }
    }
}

func RenderComposition(c *Composition, formatter *StringFormatter) {
    formatter.WriteLine("", "Composition")
    for i, p := range c.Paragraphs {
        formatter.WriteLine("", fmt.Sprintf("  Paragraph %d", i+1))
        RenderParagraph(p, formatter)
    }
}
