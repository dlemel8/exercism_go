package pangram

import (
	"strings"
)

const testVersion = 2

func IsPangram(in string) bool {
	in = strings.ToLower(in)
	letters := make([]bool, 26)

	for _, r := range in {
		if r < 'a' || r > 'z' {
			continue
		}
		letters[r-'a'] = true
	}

	for _, val := range letters {
		if !val {
			return false
		}
	}
	return true
}
