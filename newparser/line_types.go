package newparser

type LineType int

const (
    UnknownLineType LineType = -1
    LetterLineType  LineType = iota
    UpperAnnotationType
    LowerAnnotationType
    LyricLineType
)

