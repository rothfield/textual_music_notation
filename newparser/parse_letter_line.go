package newparser

type letterLineParser struct {
    tokens []Token
    pos    int
    col    int
}

func ParseLetterLine(raw string, tokens []Token) *LetterLine {
    parser := &letterLineParser{tokens: tokens}
    line := parser.parse()
    line.Raw = raw
    return line
}

func (p *letterLineParser) parse() *LetterLine {
    var elements []LetterLineElement
    for p.hasNext() {
        tok := p.peek()
        switch tok.Type {
        case Pitch, Dash:
            beat := p.parseBeat()
            elements = append(elements, *beat)
        case Barline, LeftSlur, RightSlur, Breath:
            elements = append(elements, LetterLineElement{
                Token:  p.next(),
                Column: p.col,
            })
            p.col += len(tok.Value)
        default:
            p.next()
        }
    }
    return &LetterLine{Elements: elements}
}

func (p *letterLineParser) parseBeat() *LetterLineElement {
    var sub []LetterLineElement
    startCol := p.col
    divisions := 0

    for p.hasNext() {
        tok := p.peek()
        if tok.Type != Pitch && tok.Type != Dash {
            break
        }
        sub = append(sub, LetterLineElement{
            Token:  p.next(),
            Column: p.col,
        })
        p.col += len(tok.Value)
        divisions++
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
