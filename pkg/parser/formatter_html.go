package parser

import (
	"fmt"
	"html"
	"strings"
)

func RenderHTMLComposition(c *Composition) string {
	var sb strings.Builder
	sb.WriteString(`<div class="composition">\n`)
	for _, p := range c.Paragraphs {
		sb.WriteString(RenderHTMLParagraph(p))
	}
	sb.WriteString(`</div>`)
	return sb.String()
}

func RenderHTMLParagraph(p *Paragraph) string {
	var sb strings.Builder
	sb.WriteString(`<div class="rendered-line">\n`)

	type slur struct {
		start int
		end   int
	}

	tokens := p.LetterLine.Elements
	var buffer []string
	var slurStack []int
	slurArcs := []slur{}

	// First pass: collect slur arcs and skip slur markers
	for i, el := range tokens {
		typeName := el.Token.Type
		if typeName == LeftSlur {
			slurStack = append(slurStack, i)
			continue
		} else if typeName == RightSlur && len(slurStack) > 0 {
			start := slurStack[len(slurStack)-1]
			slurStack = slurStack[:len(slurStack)-1]
			slurArcs = append(slurArcs, slur{start: start, end: i})
			continue
		}

		txt := html.EscapeString(el.Token.Value)
		htmlEl := fmt.Sprintf(`<span class="chunk">%s`, txt)
		if len(txt) > 1 {
			w := len(txt) * 12
			htmlEl += fmt.Sprintf(`<svg class="loop" width="%d" height="14"><path d="M0,0 Q%d,14 %d,0" stroke="black" fill="transparent"/></svg>`, w, w/2, w)
		}
		htmlEl += `</span>`

		buffer = append(buffer, htmlEl)
	}

	// Insert upper arcs after rendering
	for _, arc := range slurArcs {
		start := arc.start
		end := arc.end
		width := 0
		for i := start + 1; i < end; i++ {
			width += len(stripTags(buffer[i])) * 12
		}
		if width > 0 {
			path := fmt.Sprintf(`<svg class="upper-loop" width="%d" height="18"><path d="M0,18 Q%d,0 %d,18" stroke="black" fill="transparent"/></svg>`, width, width/2, width)
			buffer[start+1] = path + buffer[start+1]
		}
	}

	sb.WriteString(strings.Join(buffer, ""))
	sb.WriteString(`</div>\n`)
	return sb.String()
}

func stripTags(s string) string {
	s = strings.ReplaceAll(s, "<", "")
	s = strings.ReplaceAll(s, ">", "")
	return s
}
