package main

import (
    "regexp"
    "strings"
)

// Lexer tokenizes the letter notation into individual tokens.
func Lexer(notation string) []Token {
    var tokens []Token

    notation = strings.TrimSpace(notation)

    pitchRegex := regexp.MustCompile(`[S|r|R|g|G|m|M|P|d|D|n|N]`)
    dashRegex := regexp.MustCompile(`-`)
    barlineRegex := regexp.MustCompile(`\|:|\|\||\|`)
    slurRegex := regexp.MustCompile(`[()]`)
    breathRegex := regexp.MustCompile(`'`)
    octaveRegex := regexp.MustCompile(`[:;]`)
    mordentRegex := regexp.MustCompile(`~`)
    talaRegex := regexp.MustCompile(`[0-8+.]`)
    lyricsRegex := regexp.MustCompile(`[a-zA-Z]+`)
    spaceRegex := regexp.MustCompile(`\s+`)

    for len(notation) > 0 {
        switch {
        case pitchRegex.MatchString(notation):
            matches := pitchRegex.FindString(notation)
            tokens = append(tokens, Token{Type: Pitch, Value: matches})
            notation = notation[len(matches):]
        case dashRegex.MatchString(notation):
            matches := dashRegex.FindString(notation)
            tokens = append(tokens, Token{Type: Dash, Value: matches})
            notation = notation[len(matches):]
        case barlineRegex.MatchString(notation):
            matches := barlineRegex.FindString(notation)
            tokens = append(tokens, Token{Type: Barline, Value: matches})
            notation = notation[len(matches):]
        case slurRegex.MatchString(notation):
            matches := slurRegex.FindString(notation)
            tokens = append(tokens, Token{Type: Slur, Value: matches})
            notation = notation[len(matches):]
        case breathRegex.MatchString(notation):
            matches := breathRegex.FindString(notation)
            tokens = append(tokens, Token{Type: Breath, Value: matches})
            notation = notation[len(matches):]
        case octaveRegex.MatchString(notation):
            matches := octaveRegex.FindString(notation)
            tokens = append(tokens, Token{Type: Octave, Value: matches})
            notation = notation[len(matches):]
        case mordentRegex.MatchString(notation):
            matches := mordentRegex.FindString(notation)
            tokens = append(tokens, Token{Type: Mordent, Value: matches})
            notation = notation[len(matches):]
        case talaRegex.MatchString(notation):
            matches := talaRegex.FindString(notation)
            tokens = append(tokens, Token{Type: Tala, Value: matches})
            notation = notation[len(matches):]
        case lyricsRegex.MatchString(notation):
            matches := lyricsRegex.FindString(notation)
            tokens = append(tokens, Token{Type: Lyrics, Value: matches})
            notation = notation[len(matches):]
        case spaceRegex.MatchString(notation):
            matches := spaceRegex.FindString(notation)
            tokens = append(tokens, Token{Type: Space, Value: matches})
            notation = notation[len(matches):]
        default:
            notation = notation[1:]
        }
    }

    return tokens
}

