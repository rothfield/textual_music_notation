package parser

import (
	"strings"
)

type Paragraph struct {
	Raw              string
	LetterLine       *LetterLine
	UpperAnnotations [][]Annotation
	LowerAnnotations [][]Annotation
	LyricLines       [][]Annotation
	RawUpperLines    []string
	RawLowerLines    []string
	RawLyricLines    []string
}

func ParseParagraph(lines []string) *Paragraph {
	if len(lines) == 0 {
		return nil
	}

	Log("DEBUG", "ParseParagraph raw lines:\n%s", strings.Join(lines, "\n"))

	split := SplitLinesByType(lines)
	letter := split.LetterLine
	upperLines := split.UpperAnnotations
	lowerLines := split.LowerAnnotations
	lyricLines := split.LyricLines

	if letter == "" {
		Log("DEBUG", "ParseParagraph aborted: no letter line found.")
		return nil
	}

	Log("DEBUG", "Upper lines: %v", upperLines)
	Log("DEBUG", "Lower lines: %v", lowerLines)
	Log("DEBUG", "Lyric lines: %v", lyricLines)

	var tokens []Token
	tokens = LexLetterLine(letter)
	Log("DEBUG", "Lexed %d tokens from letter line", len(tokens))
	ParseLetterLine(letter, tokens, GuessNotation(letter))

	var (
		annotations []Annotation
		upper       []Annotation
		lower       []Annotation
		syllables   []Annotation
	)
	if len(upperLines) > 0 {

		upper = LexUpperAnnotationLine(upperLines[0])
		Log("DEBUG", "Lexed %d upper annotations", len(upper))
		annotations = append(annotations, upper...)
	}
	if len(lowerLines) > 0 {

		lower = LexLowerAnnotationLine(lowerLines[0])
		Log("DEBUG", "Lexed %d lower annotations", len(lower))
		annotations = append(annotations, lower...)
	}
	if len(lyricLines) > 0 {

		syllables = LexLyricsAnnotationLine(lyricLines[0])
		Log("DEBUG", "Lexed %d syllable annotations", len(syllables))
		annotations = append(annotations, syllables...)
	}

	var paragraph *Paragraph
	letter_line := ParseLetterLine(letter, tokens, GuessNotation(letter))
	paragraph = &Paragraph{
		Raw:              strings.Join(lines, "\n"),
		LetterLine:       letter_line,
		UpperAnnotations: [][]Annotation{upper},
		LowerAnnotations: [][]Annotation{lower},
		LyricLines:       [][]Annotation{syllables},
		RawUpperLines:    upperLines,
		RawLowerLines:    lowerLines,
		RawLyricLines:    lyricLines,
	}
	Log("DEBUG", "Calling ApplyAnnotations with %d annotations", len(annotations))
	ApplyAnnotations(paragraph, annotations)

	return paragraph
}
