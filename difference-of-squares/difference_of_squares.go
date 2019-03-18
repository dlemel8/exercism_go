package diffsquares

const testVersion = 1

func SquareOfSums(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res * res
}

func SumOfSquares(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i * i
	}
	return res
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
