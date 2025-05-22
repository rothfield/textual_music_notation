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

	if hasAnnotations(p.UpperAnnotations) {
		formatter.WriteLine("  ", "Upper Annotations")
		for _, anns := range p.UpperAnnotations {
			var tokens []string
			for _, ann := range anns {
				tokens = append(tokens, fmt.Sprintf("[%s: %s @%d]", ann.Type, ann.Value, ann.Column))
			}
			formatter.WriteLine("    ", strings.Join(tokens, " "))
		}
	}

	if hasAnnotations(p.LowerAnnotations) {
		formatter.WriteLine("  ", "Lower Annotations")
		for _, anns := range p.LowerAnnotations {
			var tokens []string
			for _, ann := range anns {
				tokens = append(tokens, fmt.Sprintf("[%s: %s @%d]", ann.Type, ann.Value, ann.Column))
			}
			formatter.WriteLine("    ", strings.Join(tokens, " "))
		}
	}

	if hasAnnotations(p.LyricLines) {
		formatter.WriteLine("  ", "Lyric Lines")
		for _, anns := range p.LyricLines {
			var tokens []string
			for _, ann := range anns {
				tokens = append(tokens, fmt.Sprintf("[%s: %s @%d]", ann.Type, ann.Value, ann.Column))
			}
			formatter.WriteLine("    ", strings.Join(tokens, " "))
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
		if el.SyllableText != "" {
			parts = append(parts, fmt.Sprintf("Syllable: %q", el.SyllableText))
		}

		formatter.WriteLine(indent, strings.Join(parts, ", "))
	}
}



func hasAnnotations(groups [][]Annotation) bool {
	for _, group := range groups {
		if len(group) > 0 {
			return true
		}
	}
	return false
}
