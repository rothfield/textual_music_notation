package parser

type LineType int

const (
	UnknownLineType LineType = -1
	PitchLineType   LineType = iota
	UpperAnnotationType
	LowerAnnotationType
	LyricLineType
)

var LineTypeNames = map[LineType]string{
	UnknownLineType:     "Unknown",
	PitchLineType:       "PitchLine",
	UpperAnnotationType: "UpperAnnotation",
	LowerAnnotationType: "LowerAnnotation",
	LyricLineType:       "LyricLine",
}
