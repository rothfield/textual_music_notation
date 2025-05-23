package newparser

func ParseLetterLine(raw string, tokens []Token) *LetterLine {
	Log("DEBUG", "ParseLetterLine: raw='%s'", raw)
	var parser *letterLineParser
    parser = &letterLineParser{tokens: tokens}
	var line *LetterLine
    line = parser.parse()
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
		var tok Token
        tok = p.peek()
		Log("DEBUG", "parse: next token = %s", tok)
		switch tok.Type {
		case Pitch, Dash:
			var beat *LetterLineElement
            beat = p.parseBeat()
			Log("DEBUG", "parse: parsed beat with %d divisions", beat.Divisions)
			elements = append(elements, *beat)
		case Barline, LeftSlur, RightSlur, Breath:
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
		var tok Token
        tok = p.peek()
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
	return t == Pitch || t == Dash || t == Breath || t == LeftSlur || t == RightSlur
}

