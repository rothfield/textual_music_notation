package main

import (
    "fmt"
    "strings"
)

// StringFormatter formats the tree structure into a string
type StringFormatter struct {
    Builder strings.Builder
}

// WriteLine writes a simple line with indentation
func (sf *StringFormatter) WriteLine(indent, text string) {
    sf.Builder.WriteString(indent + text + "\n")
}

// WriteAnnotation writes an annotation with indentation
func (sf *StringFormatter) WriteAnnotation(indent, typ, value string, column int) {
    sf.Builder.WriteString(fmt.Sprintf(indent+"  - %s: %s [Column=%d]\n", typ, value, column))
}

// WriteElement writes an element with indentation and Column position
func (sf *StringFormatter) WriteElement(indent, typ, value string, column int, octave int, mordent bool, talaMarker string, lyricText string) {
    Log("DEBUG", "Formatting Element - Type: %s, Value: %s, Column: %d", typ, value, column)
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

// WriteSubElement writes a sub-element with indentation and Column position
func (sf *StringFormatter) WriteSubElement(indent, typ, value string, column int, octave int, mordent bool, talaMarker string, lyricText string) {
    Log("DEBUG", "Formatting SubElement - Type: %s, Value: %s, Column: %d", typ, value, column)
	sf.Builder.WriteString(fmt.Sprintf(indent+"    - %s: %s [Column=%d]\n", typ, value, column))
if octave != 0 {
    sf.Builder.WriteString(fmt.Sprintf(indent+"      - Octave: %d\n", octave))
}
if mordent {
    sf.Builder.WriteString(fmt.Sprintf(indent+"      - Mordent: true\n"))
}
if talaMarker != "" {
    sf.Builder.WriteString(fmt.Sprintf(indent+"      - Tala: %s\n", talaMarker))
}
if lyricText != "" {
    sf.Builder.WriteString(fmt.Sprintf(indent+"      - Lyric: %s\n", lyricText))
}
}

