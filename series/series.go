package series

const testVersion = 2

func All(n int, s string) []string {
	size := len(s) + 1 - n
	res := make([]string, size)
	for i := 0; i < size; i++ {
		res[i] = s[i : i+n]
	}
	return res
}

func UnsafeFirst(n int, s string) string {
	first, _ := First(n, s)
	return first
}

func First(n int, s string) (first string, ok bool) {
	if len(s) < n {
		return "", false
	}
	return s[:n], true
}
