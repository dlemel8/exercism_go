package triangle

import "math"

const testVersion = 3

func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}
	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return NaT
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a != b && b != c && a != c {
		return Sca
	}
	return Iso

}

type Kind string

const NaT = "not a triangle"
const Equ = "equilateral"
const Iso = "isosceles"
const Sca = "scalene"
