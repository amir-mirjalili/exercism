package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}
	var result []string
	text2 := "And all for the want of a"
	for i := 0; i < len(rhyme)-1; i++ {
		result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}

	result = append(result, fmt.Sprintf("%s %s.", text2, rhyme[0]))
	return result
}
