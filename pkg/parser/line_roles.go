package newparser

type LineRole int

const (
    UnknownLineRole LineRole = -1
    LetterLineRole  LineRole = iota
    UpperAnnotationsLineRole
    LowerAnnotationsLineRole
    LyricLineRole
)

