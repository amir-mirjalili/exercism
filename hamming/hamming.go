package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	count := 0
	if len(a) == 0 && len(b) == 0 {
		return 0, nil
	}
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				count++
			}
		}
		return count, nil
	} else {
		return count, fmt.Errorf("not equal")
	}
}
