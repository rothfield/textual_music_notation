package newparser

func ParseLetterLine(raw string, tokens []Token) *LetterLine {
	Log("DEBUG", "ParseLetterLine: raw='%s'", raw)
	parser := &letterLineParser{tokens: tokens}
	line := parser.parse()
	line.Raw = raw
	Log("DEBUG", "ParseLetterLine: parsed %d elements", len(line.Elements))
	return line
}

type letterLineParser struct {
	tokens []Token
	pos    int
	col    int
}

func (p *letterLineParser) parse() *LetterLine {
	var elements []LetterLineElement
	for p.hasNext() {
		tok := p.peek()
		Log("DEBUG", "parse: next token = %s", tok)
		switch tok.Type {
		case PitchToken, DashToken:
			beat := p.parseBeat()
			Log("DEBUG", "parse: parsed beat with %d divisions", beat.Divisions)
			elements = append(elements, *beat)
		case BarlineToken, LeftSlur, RightSlur, BreathToken:
			elements = append(elements, LetterLineElement{
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
	return &LetterLine{Elements: elements}
}

func (p *letterLineParser) parseBeat() *LetterLineElement {
	var sub []LetterLineElement
	startCol := p.col
	divisions := 0

	Log("DEBUG", "parseBeat: starting at column %d", startCol)

	for p.hasNext() {
		tok := p.peek()
		if !isAllowedInBeat(tok.Type) {
			break
		}
		sub = append(sub, LetterLineElement{
			Token:  p.next(),
			Column: p.col,
		})
		p.col += len(tok.Value)
		divisions++
		Log("DEBUG", "parseBeat: added %s to beat at col %d", tok, p.col)
	}

	return &LetterLineElement{
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
	return t == PitchToken || t == DashToken || t == BreathToken || t == LeftSlur || t == RightSlur
}

