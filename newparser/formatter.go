package newparser

import (
	"strings"
	"fmt"
)

// StringFormatter helps build a structured and indented string output.
// It wraps strings.Builder with convenience methods for writing lines.
type StringFormatter struct {
	Builder strings.Builder
}

// WriteLine writes a single line with optional indentation.
func (f *StringFormatter) WriteLine(indent, text string) {
	f.Builder.WriteString(indent + text + "\n")
}

// WriteLines writes multiple lines, each with the same indentation.
func (f *StringFormatter) WriteLines(indent string, lines []string) {
	for _, line := range lines {
		f.WriteLine(indent, line)
	}
}

// FormatComposition displays a complete composition: raw text and paragraphs.
func FormatComposition(c *Composition, formatter *StringFormatter) {
	formatter.WriteLine("", "=== Raw Notation ===")
	formatter.WriteLines("", strings.Split(c.RawText, "\n"))
	formatter.WriteLine("", "====================")

	formatter.WriteLine("", "Composition")
	for i, p := range c.Paragraphs {
		formatter.WriteLine("", fmt.Sprintf("  Paragraph %d", i+1))
		FormatParagraph(p, formatter)
	}
}

// FormatParagraph displays the letter line and any associated annotations.
func FormatParagraph(p *Paragraph, formatter *StringFormatter) {
	FormatLetterLine(p.LetterLine, formatter, "  ")

	if len(p.UpperAnnotations) > 0 {
		formatter.WriteLine("  ", "Upper Annotations")
		for _, line := range p.UpperAnnotations {
			formatter.WriteLine("    ", line)
		}
	}

	if len(p.LowerAnnotations) > 0 {
		formatter.WriteLine("  ", "Lower Annotations")
		for _, line := range p.LowerAnnotations {
			formatter.WriteLine("    ", line)
		}
	}

	if len(p.LyricLines) > 0 {
		formatter.WriteLine("  ", "LyricLines")
		for _, line := range p.LyricLines {
			formatter.WriteLine("    ", line)
		}
	}
}

// FormatLetterLine shows the raw and parsed contents of a LetterLine.
func FormatLetterLine(l *LetterLine, formatter *StringFormatter, indent string) {
	if l == nil {
		formatter.WriteLine(indent, "LetterLine: nil")
		return
	}

	formatter.WriteLine(indent, "LetterLine")
	formatter.WriteLine(indent+"  ", "=== Raw Notation ===")
	formatter.WriteLine(indent+"  ", l.Raw)
	formatter.WriteLine(indent+"  ", "====================")

	for _, el := range l.Elements {
		writeElement(formatter, indent+"  ", el)
	}
}

// writeElement displays a single element or a beat group within the letter line.
func writeElement(formatter *StringFormatter, indent string, el LetterLineElement) {
	if el.IsBeat {
		formatter.WriteLine(indent, "- Beat:")
		for _, sub := range el.SubElements {
			writeElement(formatter, indent+"  ", sub)
		}
	} else {
		parts := []string{fmt.Sprintf("- %s: %s [Column=%d]", el.Token.Type, el.Token.Value, el.Column)}

		if el.Octave != 0 {
			parts = append(parts, fmt.Sprintf("Octave: %+d", el.Octave))
		}
		if el.Mordent {
			parts = append(parts, "Mordent: true")
		}
		if el.TalaMarker != "" {
			parts = append(parts, fmt.Sprintf("Tala: %s", el.TalaMarker))
		}
		if len(el.Syllables) > 0 {
	parts = append(parts, fmt.Sprintf("Syllables: %q", strings.Join(el.Syllables, " ")))
}

		formatter.WriteLine(indent, strings.Join(parts, ", "))
	}
}

