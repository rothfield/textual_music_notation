package main

import (
    "fmt"
    "strings"
)

type StringFormatter struct {
    Builder strings.Builder
}

func (sf *StringFormatter) WriteLine(indent, text string) {
    sf.Builder.WriteString(indent + text + "\n")
}

func (sf *StringFormatter) WriteAnnotation(indent, typ, value string, column int) {
    sf.Builder.WriteString(fmt.Sprintf(indent+"  - %s: %s [Column=%d]\n", typ, value, column))
}

func (sf *StringFormatter) WriteElement(indent, typ, value string, column int, octave int, mordent bool, talaMarker string, lyricText string) {
    sf.Builder.WriteString(fmt.Sprintf(indent+"  - %s: %s [Column=%d]\n", typ, value, column))
    if octave != 0 {
        sf.Builder.WriteString(fmt.Sprintf(indent+"    - Octave: %d\n", octave))
    }
    if mordent {
        sf.Builder.WriteString(fmt.Sprintf(indent+"    - Mordent: true\n"))
    }
    if talaMarker != "" {
        sf.Builder.WriteString(fmt.Sprintf(indent+"    - Tala: %s\n", talaMarker))
    }
    if lyricText != "" {
        sf.Builder.WriteString(fmt.Sprintf(indent+"    - Lyric: %s\n", lyricText))
    }
}
func RenderParagraph(p *Paragraph, f *StringFormatter, indent string) {
    if len(p.UpperAnnotations) > 0 {
        f.WriteLine(indent, "Upper Annotations")
        for _, ann := range p.UpperAnnotations {
            f.WriteAnnotation(indent, string(ann.Type), ann.Value, ann.Column)
        }
    }

    f.WriteLine(indent, "LetterLine")
    for _, element := range p.LetterLine.Elements {
        if element.IsBeat {
            f.WriteLine(indent, fmt.Sprintf("  - Beat: [Divisions=%d]", element.Divisions))
            for _, sub := range element.SubElements {
                f.WriteElement(indent+"    ",
                    string(sub.Token.Type),
                    sub.Token.Value,
                    sub.Column,
                    sub.Octave,
                    sub.Mordent,
                    sub.TalaMarker,
                    sub.LyricText,
                )
            }
            continue
        }

        f.WriteElement(indent,
            string(element.Token.Type),
            element.Token.Value,
            element.Column,
            element.Octave,
            element.Mordent,
            element.TalaMarker,
            element.LyricText,
        )
    }

    if len(p.LowerAnnotations) > 0 {
        f.WriteLine(indent, "Lower Annotations")
        for _, ann := range p.LowerAnnotations {
            f.WriteAnnotation(indent, string(ann.Type), ann.Value, ann.Column)
        }
    }

    if p.Lyrics != nil && len(p.Lyrics) > 0 {
        f.WriteLine(indent, "Lyrics")
        for _, ann := range p.Lyrics {
            f.WriteAnnotation(indent, string(ann.Type), ann.Value, ann.Column)
        }
    }
}

func zRenderParagraph(p *Paragraph, f *StringFormatter, indent string) {
    if len(p.UpperAnnotations) > 0 {
        f.WriteLine(indent, "Upper Annotations")
        for _, ann := range p.UpperAnnotations {
            f.WriteAnnotation(indent, string(ann.Type), ann.Value, ann.Column)
        }
    }

    f.WriteLine(indent, "LetterLine")
    for _, element := range p.LetterLine.Elements {
        typ := element.Token.Type
        if typ == Pitch || typ == Dash || typ == Barline || typ == Breath {
            f.WriteElement(indent,
                string(typ),
                element.Token.Value,
                element.Token.Column,
                element.Octave,
                element.Mordent,
                element.TalaMarker,
                element.LyricText,
            )
        }
    }

    if len(p.LowerAnnotations) > 0 {
        f.WriteLine(indent, "Lower Annotations")
        for _, ann := range p.LowerAnnotations {
            f.WriteAnnotation(indent, string(ann.Type), ann.Value, ann.Column)
        }
    }

    if p.Lyrics != nil && len(p.Lyrics) > 0 {
        f.WriteLine(indent, "Lyrics")
        for _, ann := range p.Lyrics {
            f.WriteAnnotation(indent, string(ann.Type), ann.Value, ann.Column)
        }
    }
}
func DisplayCompositionTree(c *Composition, f *StringFormatter) {
    f.WriteLine("", "Composition")
    for i, p := range c.Paragraphs {
        f.WriteLine("", fmt.Sprintf("  Paragraph %d", i+1))
        RenderParagraph(&p, f, "    ")
    }
}

