package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

func IsIsogram(word string) bool {
	alreadySeen := make(map[rune]bool)
	for _, r := range strings.ToLower(word) {
		if !unicode.IsLetter(r) {
			continue
		}
		if _, ok := alreadySeen[r]; ok {
			return false
		}
		alreadySeen[r] = true
	}
	return true
}
