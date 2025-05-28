package parser

import (
	"fmt"
	"strings"
)

// StringFormatter helps build a structured and indented string output.
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
	formatter.WriteLine("  ", fmt.Sprintf("Notation: %s", p.Notation))
	FormatLine(p.Line, formatter, "  ")

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

// FormatLine shows the raw and parsed contents of a Line.
func FormatLine(l *Line, formatter *StringFormatter, indent string) {
	if l == nil {
		formatter.WriteLine(indent, "Line: nil")
		return
	}

	formatter.WriteLine(indent, "Line")
	formatter.WriteLine(indent+"  ", "=== Raw Notation ===")
	formatter.WriteLine(indent+"  ", l.Raw)
	formatter.WriteLine(indent+"  ", "====================")

	for _, el := range l.Elements {
		writeElement(formatter, indent+"  ", el)
	}
}

// writeElement displays a single element or a beat group within the letter line.
func writeElement(formatter *StringFormatter, indent string, el Element) {
	if el.IsBeat {
		formatter.WriteLine(indent, "- Beat:")
		for _, sub := range el.SubElements {
			writeElement(formatter, indent+"  ", sub)
		}
		return
	}

	col := el.Column
	val := el.Token.Value

	switch el.Token.Type {
	case "Pitch":
		formatter.WriteLine(indent, fmt.Sprintf("- Pitch: %s [Column=%d], Octave: %d", val, col, el.Octave))
		if el.Mordent {
			formatter.WriteLine(indent+"  ", "Mordent: true")
		}
		if el.Tala != "" {
			formatter.WriteLine(indent+"  ", fmt.Sprintf("Tala: %s", el.Tala))
		}
		if el.Syllable != "" {
			formatter.WriteLine(indent+"  ", fmt.Sprintf("Syllable: %q", el.Syllable))
		}
		if len(el.ExtraSyllables) > 0 {
			formatter.WriteLine(indent+"  ", fmt.Sprintf("ExtraSyllables: %q", strings.Join(el.ExtraSyllables, " ")))
		}
	default:
		formatter.WriteLine(indent, fmt.Sprintf("- %s: %q [Column=%d]", el.Token.Type.String(), val, col))
	}
}

// hasAnnotations checks whether any annotation lines are non-empty.
func hasAnnotations(groups [][]Annotation) bool {
	for _, group := range groups {
		if len(group) > 0 {
			return true
		}
	}
	return false
}

