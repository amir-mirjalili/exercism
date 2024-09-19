package triangle

type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if a < 0 || b < 0 || c < 0 || (a == 0 && b == c && c == 0) || (a+b < c || a+c < b || b+c < a) {
		return NaT
	} else if a == b && a == c && b == c {
		return Equ
	} else if (a == b && a != b) || (a != c && b == c) || (b != c && b == a) || (a == c && a != b) {
		return Iso
	} else {
		return Sca
	}
}
