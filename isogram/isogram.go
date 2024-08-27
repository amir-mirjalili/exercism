package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	wordArr := strings.Split(word, "")
	for index, i := range wordArr {
		for index2, i2 := range wordArr {
			if index != index2 && i == i2 && i != " " && i != "-" && i != "_" {
				return false
			}
		}
	}
	return true
}
