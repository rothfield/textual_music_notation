package parser

type LineRole int

const (
	UnknownLineRole LineRole = -1
	PitchLineRole   LineRole = iota
	UpperAnnotationsLineRole
	LowerAnnotationsLineRole
	LyricLineRole
)
