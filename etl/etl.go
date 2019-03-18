package etl

import "strings"

const testVersion = 1

func Transform(in map[int][]string) map[string]int {
	res := make(map[string]int)
	for key, val := range in {
		for _, c := range val {
			res[strings.ToLower(c)] = key
		}
	}
	return res
}
