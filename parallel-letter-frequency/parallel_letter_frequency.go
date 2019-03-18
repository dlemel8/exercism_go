package letter

const testVersion = 1

func ConcurrentFrequency(input []string) FreqMap {
	res := make(FreqMap)
	c := make(chan FreqMap)
	for _, s := range input {
		go func(str string) {
			c <- Frequency(str)
		}(s)
	}
	for range input {
		for r, count := range <-c {
			res[r] += count
		}
	}
	return res
}
