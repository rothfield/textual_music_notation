package main

type letterLineParser struct {
    tokens []Token
    pos    int
    col    int
}

func ParseLetterLine(tokens []Token) *LetterLine {
    parser := &letterLineParser{tokens: tokens}
    return parser.parse()
}

func (p *letterLineParser) parse() *LetterLine {
    var elements []LetterLineElement

    for p.hasNext() {
        tok := p.peek()
        switch tok.Type {
        case Pitch, Dash:
            beat := p.parseBeat()
            Log("DEBUG", "**parseBeat retruns, beat=%s", beat)

            elements = append(elements, *beat)
        case Barline:
            elements = append(elements, LetterLineElement{
                Token:  p.next(),
                Column: p.col,
            })
            p.col += len(tok.Value)
        default:
            Log("WARN", "Unhandled token type: %s", tok.Type)
            p.next()
        }
    }

    return &LetterLine{Elements: elements}
}

func (p *letterLineParser) parseBeat() *LetterLineElement {
    startCol := p.col
    var sub []LetterLineElement
    divisions := 0

    for p.hasNext() {
        tok := p.peek()
        Log("DEBUG", "**In parseBeat, tok=%v", tok)
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
    if p.hasNext() {
        return p.tokens[p.pos]
    }
    return Token{}
}

func (p *letterLineParser) next() Token {
    tok := p.peek()
    p.pos++
    return tok
}

type LetterLineElement struct {
    Token       Token
    Column      int
    SubElements []LetterLineElement
    IsBeat      bool
    Divisions   int // number of subdivisions if IsBeat is true
    Octave      int
    Mordent     bool
    TalaMarker  string
    LyricText   string
}

type LetterLine struct {
    Elements []LetterLineElement
}

