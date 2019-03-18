package pascal

const testVersion = 1

func newLayer(prev []int) []int {
	size := len(prev) + 1
	res := make([]int, size)
	res[0], res[size-1] = 1, 1
	for i := 1; i < size-1; i++ {
		res[i] = prev[i-1] + prev[i]
	}
	return res
}

func Triangle(n int) [][]int {
	res := make([][]int, n)
	res[0] = []int{1}
	for i := 1; i < n; i++ {
		res[i] = newLayer(res[i-1])
	}
	return res
}
