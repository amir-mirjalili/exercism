package dna

import (
	"errors"
	"strings"
)

type Histogram map[rune]uint

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	dnaArray := DNA(strings.ToUpper(string(d)))

	for _, dna := range dnaArray {
		switch dna {
		case 'A', 'C', 'G', 'T':
			h[dna]++
		default:
			return Histogram{}, errors.New("invalid nucleotide")
		}
	}
	return h, nil
}
