package luhn

import (
	"regexp"
	"strings"
)

const testVersion = 2

var nonValidChars = regexp.MustCompile("[^ 0-9]+")
var spacesReplacer = strings.NewReplacer(" ", "")

func Valid(input string) bool {
	if nonValidChars.MatchString(input) {
		return false
	}

	digits := spacesReplacer.Replace(input)
	if len(digits) < 2 {
		return false
	}

	res := 0
	digitsParity := (len(digits) - 1) % 2
	for i := len(digits) - 1; i >= 0; i-- {
		digit := int(digits[i] - '0')
		if i%2 != digitsParity && digit != 9 {
			res += (digit * 2) % 9
		} else {
			res += digit
		}
	}
	return res%10 == 0
}
