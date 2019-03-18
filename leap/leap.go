package leap

import "math"

const testVersion = 3

func isDivided(number1 int, number2 float64) bool {
	return math.Mod(float64(number1), number2) == 0
}

func IsLeapYear(year int) bool {
	return isDivided(year, 400) || (isDivided(year, 4) && !isDivided(year, 100))
}
