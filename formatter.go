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
func (sf *StringFormatter) WriteAnnotation(indent, typ, value string) {
    sf.Builder.WriteString(fmt.Sprintf(indent+"  - %s: %s\n", typ, value))
}

// WriteElement writes an element with indentation and X position
func (sf *StringFormatter) WriteElement(indent, typ, value string, x int) {
    sf.Builder.WriteString(fmt.Sprintf(indent+"  - %s: %s [X=%d]\n", typ, value, x))
}

// WriteSubElement writes a sub-element with indentation and X position
func (sf *StringFormatter) WriteSubElement(indent, typ, value string, x int) {
    sf.Builder.WriteString(fmt.Sprintf(indent+"    - %s: %s [X=%d]\n", typ, value, x))
}

