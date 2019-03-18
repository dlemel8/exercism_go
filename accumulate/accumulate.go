package accumulate

const testVersion = 1

func Accumulate(in []string, f func(string) string) []string {
	var res = make([]string, len(in))
	for i := range in {
		res[i] = f(in[i])
	}
	return res
}
