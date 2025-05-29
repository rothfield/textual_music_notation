package parser

import (
	"fmt"
	"html"
	"strings"
)

func RenderHTMLComposition(c *Composition) string {
	var sb strings.Builder
	sb.WriteString(`<composition>`) // semantic root tag
	for _, p := range c.Paragraphs {
		sb.WriteString(RenderHTMLParagraph(p))
	}
	sb.WriteString(`</composition>`)
	return sb.String()
}

func RenderHTMLParagraph(p *Paragraph) string {
	var sb strings.Builder
	sb.WriteString(`<paragraph>`) // custom tag for semantic grouping

	inSlur := false
	var slurGroup []string

	flushSlur := func() {
		if len(slurGroup) == 0 {
			return
		}
		groupHTML := strings.Join(slurGroup, "")
		sb.WriteString(`<slur-group>`) // semantic grouping for slurs
		sb.WriteString(groupHTML)
		sb.WriteString(`</slur-group>`)
		slurGroup = nil
	}

	for _, el := range p.Line.Elements {
		switch el.Token.Type {
		case TokenTypeLeftSlur:
			flushSlur()
			inSlur = true
			continue
		case TokenTypeRightSlur:
			inSlur = false
			flushSlur()
			continue
		}

		htmlChunk := renderElementHTML(&el)
		if inSlur {
			slurGroup = append(slurGroup, htmlChunk)
		} else {
			sb.WriteString(htmlChunk)
		}
	}

	flushSlur()
	sb.WriteString(`</paragraph>`)
	return sb.String()
}

func renderElementHTML(el *Element) string {
	switch el.Token.Type {
	case TokenTypeBarline:
		return fmt.Sprintf(`<barline>%s</barline>`, html.EscapeString(el.Token.Value))
	case TokenTypeBreath:
		return `<breath>'</breath>`
	case TokenTypeDash:
		return `<dash>&mdash;</dash>`
	case TokenTypePitch:
		return renderNoteHTML(el)
	default:
		if len(el.SubElements) > 0 {
			return renderBeatHTML(el)
		}
		return fmt.Sprintf(`<%s>%s</%s>`, strings.ToLower(el.Token.Type.String()), html.EscapeString(el.Token.Value), strings.ToLower(el.Token.Type.String()))
	}
}

func renderNoteHTML(el *Element) string {
	var sb strings.Builder
	sb.WriteString(`<note>`) // custom tag

	sb.WriteString(fmt.Sprintf(`<pitch>%s</pitch>`, html.EscapeString(el.Token.Value)))

	if el.Octave != 0 {
		sb.WriteString(fmt.Sprintf(`<octave data-octave="%d">%s</octave>`, el.Octave, "."))
	}
	if el.Syllable != "" {
		sb.WriteString(fmt.Sprintf(`<syllable>%s</syllable>`, html.EscapeString(el.Syllable)))
	}

	sb.WriteString(`</note>`)
	return sb.String()
}

func renderBeatHTML(el *Element) string {
	var sb strings.Builder

	// Determine divisions (default is 4, omit if 4)
	divs := el.Divisions
	if divs != 4 && divs > 0 {
		sb.WriteString(fmt.Sprintf(`<beat data-divisions="%d">`, divs))
	} else {
		sb.WriteString(`<beat>`)
	}

	for i := range el.SubElements {
		sub := renderElementHTML(&el.SubElements[i])
		sb.WriteString(sub)
	}

	sb.WriteString(`</beat>`)
	return sb.String()
}
