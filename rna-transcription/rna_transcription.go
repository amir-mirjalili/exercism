package strand

import "strings"

func ToRNA(dna string) string {
	words := map[string]string{"G": "C", "C": "G", "T": "A", "A": "U"}
	dnaArray := strings.Split(dna, "")
	result := ""
	for _, dn := range dnaArray {
		switch dn {
		case "G":
			result += words["G"]
		case "C":
			result += words["C"]
		case "T":
			result += words["T"]
		case "A":
			result += words["A"]
		case "U":
			result += words["U"]
		default:
			result += words[dn]
		}
	}
	return result
}
