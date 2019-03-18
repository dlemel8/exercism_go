package reverse

import "unicode/utf8"

func String(s string) string {
	size := utf8.RuneCountInString(s)
	runes := []rune(s)
	for i := 0; i < size/2; i++ {
		runes[i], runes[size-1-i] = runes[size-1-i], runes[i]
	}
	return string(runes)
}
