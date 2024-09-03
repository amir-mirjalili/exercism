package romannumerals

import (
	"errors"
	"fmt"
)

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", errors.New(fmt.Sprintf("%d_is_out_of_range", input))
	}
	var numbers = []struct {
		Value  int
		Symbol string
	}{
		{Value: 1000, Symbol: "M"},
		{Value: 900, Symbol: "CM"},
		{Value: 500, Symbol: "D"},
		{Value: 400, Symbol: "CD"},
		{Value: 100, Symbol: "C"},
		{Value: 90, Symbol: "XC"},
		{Value: 50, Symbol: "L"},
		{Value: 40, Symbol: "XL"},
		{Value: 10, Symbol: "X"},
		{Value: 9, Symbol: "IX"},
		{Value: 5, Symbol: "V"},
		{Value: 4, Symbol: "IV"},
		{Value: 1, Symbol: "I"},
	}
	result := ""
	for _, val := range numbers {
		for input >= val.Value {
			result += val.Symbol
			input -= val.Value
		}

	}
	return result, nil
}
