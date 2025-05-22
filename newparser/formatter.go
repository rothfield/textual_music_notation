package newparser

import (
	"strings"
	"fmt"
)

type StringFormatter struct {
	Builder strings.Builder
}

func (f *StringFormatter) WriteLine(indent, text string) {
	f.Builder.WriteString(indent + text + "\n")
}

func (f *StringFormatter) WriteLines(indent string, lines []string) {
	for _, line := range lines {
		f.WriteLine(indent, line)
	}
}

func RenderComposition(c *Composition, formatter *StringFormatter) {
	formatter.WriteLine("", "=== Raw Notation ===")
	formatter.WriteLines("", strings.Split(c.RawText, "\n"))
	formatter.WriteLine("", "====================")

	formatter.WriteLine("", "Composition")
	for i, p := range c.Paragraphs {
		formatter.WriteLine("", fmt.Sprintf("  Paragraph %d", i+1))
		RenderParagraph(p, formatter)
	}
}

func RenderParagraph(p *Paragraph, formatter *StringFormatter) {
	RenderLetterLine(p.LetterLine, formatter, "  ")

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

	if len(p.Lyrics) > 0 {
		formatter.WriteLine("  ", "Lyrics")
		for _, line := range p.Lyrics {
			formatter.WriteLine("    ", line)
		}
	}
}

func RenderLetterLine(l *LetterLine, formatter *StringFormatter, indent string) {
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

func writeElement(formatter *StringFormatter, indent string, el LetterLineElement) {
	if el.IsBeat {
		formatter.WriteLine(indent, "- Beat:")
		for _, sub := range el.SubElements {
			writeElement(formatter, indent+"  ", sub)
		}
	} else {
		s := fmt.Sprintf("- %s: %s [Column=%d]", el.Token.Type, el.Token.Value, el.Column)
		if el.Octave != 0 {
			s += fmt.Sprintf(", Octave: %+d", el.Octave)
		}
		if el.Mordent {
			s += ", Mordent: true"
		}
		if el.TalaMarker != "" {
			s += fmt.Sprintf(", Tala: %s", el.TalaMarker)
		}
		if el.SyllableText != "" {
			s += fmt.Sprintf(", Lyric: %q", el.SyllableText)
		}
		formatter.WriteLine(indent, s)
	}
}

