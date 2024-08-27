package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	var result []string
	if number%3 == 0 {
		result = append(result, "Pling")
	}
	if number%5 == 0 {
		result = append(result, "Plang")
	}
	if number%7 == 0 {
		result = append(result, "Plong")
	}
	if len(result) > 0 {
		return strings.Join(result, "")
	} else {
		return strconv.Itoa(number)
	}
}
