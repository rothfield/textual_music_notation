package newparser

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

func RenderLetterLine(line *LetterLine, f *StringFormatter, indent string) {
    f.WriteLine(indent, "LetterLine")
    for _, element := range line.Elements {
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
        } else {
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
    }
}
