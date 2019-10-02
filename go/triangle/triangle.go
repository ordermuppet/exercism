package triangle

import "math"

// Kind represents a kind of triangle
type Kind int

// kind of triangle (equilateral, isosceles, scalene, or invalid)
const (
	NaT = 0 // not a triangle
	Equ = 1 // equilateral
	Iso = 2 // isosceles
	Sca = 3 // scalene
)

// KindFromSides returns which kind of triangle (equilateral, isosceles, or scalene) these sides represent
func KindFromSides(a, b, c float64) Kind {
	// For a shape to be a triangle at all, all sides have to be of length > 0, but not infinite,
	if a <= 0 || b <= 0 || c <= 0 || math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return NaT
	}

	// and the sum of the lengths of any two sides must be greater than or equal to the length of the third side.
	if !((a+b >= c) && (a+c >= b) && (b+c >= a)) {
		return NaT
	}

	if a == b && a == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}
