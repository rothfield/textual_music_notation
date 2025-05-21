package main

import (
    "fmt"
    "textual_music_notation/newparser"  // ✅ or your actual module path
)

func DebugLetterLine() {
    Log("DEBUG", "In DebugLetterLine")

    raw := "S(-r)|P"
    tokens := newparser.LexLetterLine(raw)
    line := newparser.ParseLetterLine(tokens)

    formatter := &newparser.StringFormatter{}
    newparser.RenderLetterLine(line, formatter, "")

    fmt.Println("==== Formatter Test Output ====")
    fmt.Println(formatter.Builder.String())
}

