package acronym

import (
	"bytes"
	"strings"
)

const testVersion = 3

func split(r rune) bool {
	return r == ' ' || r == '-'
}
func Abbreviate(in string) string {
	var res bytes.Buffer

	for _, word := range strings.FieldsFunc(in, split) {
		res.WriteString(strings.ToUpper(string(word[0])))
	}

	return res.String()

}
