package summultiples

const testVersion = 2

func SumMultiples(limit int, divisors ...int) int {
	numbers := make([]bool, limit)
	for _, num := range divisors {
		for i := num; i < limit; i += num {
			numbers[i] = true
		}
	}
	res := 0
	for i := range numbers {
		if numbers[i] {
			res += i
		}
	}
	return res
}
