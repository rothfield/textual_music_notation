package main

import (
    "fmt"
    "strings"
)

var noteMap = map[string]string{
    "S": "c",
    "r": "db",
    "R": "d",
    "g": "eb",
    "G": "e",
    "m": "f",
    "M": "fis",
    "P": "g",
    "d": "ab",
    "D": "a",
    "n": "bb",
    "N": "b",
}

// ✅ GenerateLilyPond converts a parsed paragraph into LilyPond notation.
func GenerateLilyPond(paragraph Paragraph) string {
    lilypond := "\\version \"2.24.0\"\n"
    lilypond += "\\relative c' {\n"

    for _, element := range paragraph.LetterLine.Elements {
        if element.Token.Type == "Pitch" {
            pitch := noteMap[element.Token.Value]

            // Octave adjustment
            if element.Octave > 0 {
                pitch += strings.Repeat("'", element.Octave)
            } else if element.Octave < 0 {
                pitch += strings.Repeat(",", -element.Octave)
            }

            // Duration Calculation
            duration := calculateDuration(element)

            // Add annotations
            if element.Mordent {
                pitch = fmt.Sprintf("\\mordent %s", pitch)
            }
            if element.TalaMarker != "" {
                lilypond += fmt.Sprintf("\\markup { \"Tala %s\" }\n", element.TalaMarker)
            }

            // Append the pitch and duration to the line
            lilypond += fmt.Sprintf("  %s%s ", pitch, duration)

            // Lyrics attachment
            if element.LyricText != "" {
                lilypond += fmt.Sprintf("\\lyricmode { %s } ", element.LyricText)
            }
        }
    }
    
    lilypond += "\n}\n"
    return lilypond
}

// ✅ calculateDuration converts dashes to fractional beats
func calculateDuration(element LetterLineElement) string {
    dashCount := strings.Count(element.Token.Value, "-")
    totalDuration := 1 << dashCount // 2^dashCount

    if totalDuration == 1 {
        return "4" // Quarter note
    } else if totalDuration == 2 {
        return "8"
    } else if totalDuration == 4 {
        return "16"
    } else {
        // If it's not a power of 2, represent as a tuplet
        return fmt.Sprintf("\\tuplet %d/2 { %s }", totalDuration, element.Token.Value)
    }
}

