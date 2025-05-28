package parser

import (
	"fmt"
	"html"
	"strings"
)

func RenderHTMLComposition(c *Composition) string {
	var sb strings.Builder
	sb.WriteString(`<div class="composition">`)
	for _, p := range c.Paragraphs {
		sb.WriteString(RenderHTMLParagraph(p))
	}
	sb.WriteString(`</div>`)
	return sb.String()
}

func RenderHTMLParagraph(p *Paragraph) string {
	var sb strings.Builder
	sb.WriteString(`<div class="paragraph">`)

	inSlur := false
	var slurGroup []string

	flushSlur := func() {
		if len(slurGroup) == 0 {
			return
		}
		groupHTML := strings.Join(slurGroup, "")
		sb.WriteString(`<div class="slur-group">`)
		sb.WriteString(groupHTML)
		sb.WriteString(`</div>`)
		slurGroup = nil
	}

	for _, el := range p.Line.Elements {
		switch el.Token.Type {
		case LeftSlur:
			flushSlur()
			inSlur = true
			continue
		case RightSlur:
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
	sb.WriteString(`</div>`)
	return sb.String()
}

func renderElementHTML(el *Element) string {
	txt := html.EscapeString(el.Token.Value)
	typ := strings.ToLower(el.Token.Type.String())

	if len(el.SubElements) > 0 {
		var sb strings.Builder
		sb.WriteString(`<span class="beat">`)
		for i := range el.SubElements {
			sub := renderElementHTML(&el.SubElements[i])
			sb.WriteString(sub)
		}
		sb.WriteString(`</span>`)
		return sb.String()
	}

	var sb strings.Builder
	sb.WriteString(`<span class="pitch-wrapper">`)

	// Upper octave: • or :
	if el.Octave > 0 {
		symbol := "•"
		if el.Octave == 2 {
			symbol = ":"
		}
		sb.WriteString(fmt.Sprintf(`<span class="upper">%s</span>`, symbol))
	}

	sb.WriteString(fmt.Sprintf(`<span class="pitch %s">%s</span>`, typ, txt))

	// Lower octave: • or :
	if el.Octave < 0 {
		symbol := "•"
		if el.Octave == -2 {
			symbol = ":"
		}
		sb.WriteString(fmt.Sprintf(`<span class="lower octave">%s</span>`, symbol))
	}

	// Lyric
	if el.Syllable != "" {
		sb.WriteString(`<span class="lyric-container">`)
		sb.WriteString(fmt.Sprintf(`<span class="lower lyric">%s</span>`, html.EscapeString(el.Syllable)))
		sb.WriteString(`</span>`)
	}

	sb.WriteString(`</span>`)
	return sb.String()
}
