package parser

import (
	"fmt"
	"html"
	"strings"
)

// Public function to render Composition to HTML
func CompositionToHTML(c *Composition) string {
	var sb strings.Builder
	sb.WriteString(`<composition>`)
	for _, p := range c.Paragraphs {
		sb.WriteString(ParagraphToHTML(p))
	}
	sb.WriteString(`</composition>`)
	return sb.String()
}

// Public function to render Paragraph to HTML
func ParagraphToHTML(p *Paragraph) string {
	var sb strings.Builder
	sb.WriteString(`<paragraph>`)

	for _, el := range p.Line.Elements {
		sb.WriteString(RenderElementToHTML(&el))
	}

	sb.WriteString(`</paragraph>`)
	return sb.String()
}

// Public function to render Note to HTML
func NoteToHTML(el *Element) string {
	var sb strings.Builder
	sb.WriteString(`<note>`)
	sb.WriteString(fmt.Sprintf(`<pitch>%s</pitch>`, html.EscapeString(el.Token.Value)))

	// Handle octave, mordent, syllables
	if el.Octave != 0 {
		sb.WriteString(fmt.Sprintf(`<octave data-octave="%d">`, el.Octave))
		if el.Octave == 1 || el.Octave == -1 {
			sb.WriteString(`•`)
		} else if el.Octave == 2 || el.Octave == -2 {
			sb.WriteString(`:`)
		}
		sb.WriteString(`</octave>`)
	}

	if el.Mordent {
		sb.WriteString(`<mordent>~</mordent>`)
	}
	if el.Tala != "" {
		sb.WriteString(fmt.Sprintf(`<tala>%s</tala>`, html.EscapeString(el.Tala)))
	}
	if el.Syllable != "" {
		sb.WriteString(fmt.Sprintf(`<syllable>%s</syllable>`, html.EscapeString(el.Syllable)))
	}
	for _, syl := range el.ExtraSyllables {
		sb.WriteString(fmt.Sprintf(`<syllable>%s</syllable>`, html.EscapeString(syl)))
	}

	sb.WriteString(`</note>`)
	return sb.String()
}

// Public function to render Beat to HTML
func BeatToHTML(el *Element) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(`<beat data-divisions="%d">`, len(el.SubElements)))
	for i := range el.SubElements {
		sb.WriteString(RenderElementToHTML(&el.SubElements[i]))
	}
	sb.WriteString(`</beat>`)
	return sb.String()
}

// Public function to handle LeftSlur rendering
func HandleLeftSlur(el *Element) string {
	if el.Token.Type != TokenTypeLeftSlur {
		panic(fmt.Sprintf("HandleLeftSlur called on invalid token type: %v", el.Token.Type))
	}
	return `<slur>`
}

// Public function to handle RightSlur rendering
func HandleRightSlur(el *Element) string {
	if el.Token.Type != TokenTypeRightSlur {
		panic(fmt.Sprintf("HandleRightSlur called on invalid token type: %v", el.Token.Type))
	}
	return `</slur>`
}

// Main render function for each element type
func RenderElementToHTML(el *Element) string {
	switch el.Token.Type {
	case TokenTypeLeftSlur:
		return HandleLeftSlur(el)
	case TokenTypeRightSlur:
		return HandleRightSlur(el)
	case TokenTypeBarline:
		return fmt.Sprintf(`<barline>%s</barline>`, html.EscapeString(el.Token.Value))
	case TokenTypeBreath:
		return `<breath>'</breath>`
	case TokenTypeDash:
		return `<dash>&mdash;</dash>`
	case TokenTypePitch:
		return NoteToHTML(el)
	default:
		if len(el.SubElements) > 0 {
			return BeatToHTML(el)
		}
		return fmt.Sprintf(
			`<%s>%s</%s>`,
			strings.ToLower(el.Token.Type.String()),
			html.EscapeString(el.Token.Value),
			strings.ToLower(el.Token.Type.String()),
		)
	}
}

