package newparser

type LineType int

const (
    LetterLineType LineType = iota
    UpperAnnotationType
    LowerAnnotationType
    SyllableType
)
