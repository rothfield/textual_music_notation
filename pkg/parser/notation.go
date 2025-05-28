package parser

import ()

type Notation uint8

const (
	Western Notation = iota
	Number
	Sargam
)

func (n Notation) String() string {
	switch n {
	case Western:
		return "Western"
	case Number:
		return "Number"
	case Sargam:
		return "Sargam"
	default:
		return "Unknown"
	}
}

type PitchCode struct {
	DiatonicNote uint8 // 0 = C, 1 = D, ..., 6 = B
	Alteration   int8  // -2 (double flat) to +2 (double sharp)
}

var pitchCodes = map[Notation]map[string]PitchCode{
	Western: {
		"C": {0, 0}, "Db": {1, -1},
		"D": {1, 0}, "Eb": {2, -1},
		"E": {2, 0}, "F": {3, 0},
		"F#": {3, 1}, "G": {4, 0},
		"Ab": {5, -1}, "A": {5, 0},
		"Bb": {6, -1}, "B": {6, 0},
	},
	Number: {
		"1":  {0, 0},
		"2b": {1, -1},
		"2":  {1, 0},
		"3b": {2, -1},
		"3":  {2, 0},
		"4":  {3, 0},
		"5":  {4, 0},
		"6":  {5, 0},
		"7":  {6, 0},
	},
	Sargam: {
		"S": {0, 0},
		"r": {1, -1},
		"R": {1, 0},
		"g": {2, -1},
		"G": {2, 0},
		"m": {3, 0},
		"M": {3, 1},
		"P": {4, 0},
		"d": {5, -1},
		"D": {5, 0},
		"n": {6, -1},
		"N": {6, 0},
	},
}

func LookupPitch(symbol string, system Notation) (PitchCode, bool) {
	if sysMap, ok := pitchCodes[system]; ok {
		spelling, found := sysMap[symbol]
		return spelling, found
	}
	return PitchCode{}, false
}

func GuessNotation(line string) Notation {
	scores := map[Notation]int{
		Western: 0,
		Number:  0,
		Sargam:  0,
	}

	lexers := map[Notation]func(string) []Token{
		Western: LexABC,
		Number:  LexLineNumber,
		Sargam:  LexLineSargam,
	}

	for system, lexer := range lexers {
		tokens := lexer(line)
		for _, tok := range tokens {
			if tok.Type == Pitch {
				scores[system]++
			}
		}
	}

	best := Western
	max := -1
	for system, score := range scores {
		if score > max {
			best = system
			max = score
		}
	}

	return best
}
