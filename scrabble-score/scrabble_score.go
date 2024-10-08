package scrabble

import "strings"

func Score(word string) int {
	word = strings.ToUpper(word)
	wordArr := strings.Split(word, "")
	points := map[string]int{
		"A": 1, "U": 1,
		"O": 1, "I": 1,
		"E": 1, "L": 1,
		"N": 1, "S": 1,
		"R": 1, "T": 1,
		"D": 2, "G": 2,
		"B": 3, "C": 3,
		"P": 3, "M": 3,
		"F": 4, "H": 4,
		"V": 4, "W": 4,
		"Y": 4, "K": 5,
		"J": 8, "X": 8,
		"Q": 10, "Z": 10,
	}
	score := 0
	for _, i := range wordArr {
		for index, point := range points {
			if i == index {
				score += point
			}
		}
	}
	return score
}
