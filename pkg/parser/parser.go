package parser

func ParseLine(raw string, tokens []Token, system Notation) *Line {
	Log("DEBUG", "ParseLine: raw='%s'", raw)
	var parser *letterLineParser
	parser = &letterLineParser{tokens: tokens, system: system}
	var line *Line
	line = parser.parse()
	line.Raw = raw
	Log("DEBUG", "ParseLine: parsed %d elements", len(line.Elements))
	return line
}

type letterLineParser struct {
	tokens []Token
	system Notation
	pos    int
	col    int
}

func (p *letterLineParser) parse() *Line {
	var elements []Element
	for p.hasNext() {
		var tok Token
		tok = p.peek()
		Log("DEBUG", "parse: next token = %s", tok)
		switch tok.Type {
		case Pitch, Dash:
			var beat *Element
			beat = p.parseBeat()
			Log("DEBUG", "parse: parsed beat with %d divisions", beat.Divisions)
			elements = append(elements, *beat)
		case Barline, LeftSlur, RightSlur, Breath:
			elements = append(elements, Element{
				Token:  p.next(),
				Column: p.col,
			})
			p.col += len(tok.Value)
			Log("DEBUG", "parse: added top-level token %s at column %d", tok, p.col)
		default:
			Log("DEBUG", "parse: skipping unknown token %s", tok)
			p.next()
		}
	}
	return &Line{Elements: elements}
}

func (p *letterLineParser) parseBeat() *Element {
	var sub []Element
	startCol := p.col
	divisions := 0

	Log("DEBUG", "parseBeat: starting at column %d", startCol)

	for p.hasNext() {
		var tok Token
		tok = p.peek()
		if !isAllowedInBeat(tok.Type) {
			break
		}
		sub = append(sub, Element{
			Token:  p.next(),
			Column: p.col,
		})
		p.col += len(tok.Value)
		divisions++
		Log("DEBUG", "parseBeat: added %s to beat at col %d", tok, p.col)
	}

	return &Element{
		IsBeat:      true,
		Column:      startCol,
		SubElements: sub,
		Divisions:   divisions,
	}
}

func (p *letterLineParser) hasNext() bool {
	return p.pos < len(p.tokens)
}

func (p *letterLineParser) peek() Token {
	return p.tokens[p.pos]
}

func (p *letterLineParser) next() Token {
	tok := p.tokens[p.pos]
	p.pos++
	return tok
}

func isAllowedInBeat(t TokenType) bool {
	return t == Pitch || t == Dash || t == Breath || t == LeftSlur || t == RightSlur
}
