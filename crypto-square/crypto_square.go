package cryptosquare

import (
	"bytes"
	"math"
	"regexp"
	"strings"
)

const testVersion = 2

var alphanumeric = regexp.MustCompile("[^a-z0-9]+")

func Encode(plain string) string {
	if len(plain) == 0 {
		return ""
	}
	normalized := normalize(plain)
	_, c := calcRectangle(normalized)
	return reorderByColumns(normalized, c)
}

func normalize(s string) string {
	return alphanumeric.ReplaceAllString(strings.ToLower(s), "")
}

func calcRectangle(s string) (int, int) {
	size := float64(len(s))
	r := math.Sqrt(size)
	c := math.Ceil(size / r)
	return int(r), int(c)
}

func reorderByColumns(s string, columnsNum int) string {
	parts := make([]string, columnsNum)
	for col := 0; col < columnsNum; col++ {
		var buffer bytes.Buffer
		for pos := 0; pos+col < len(s); pos += columnsNum {
			buffer.WriteByte(s[pos+col])
		}
		parts[col] = buffer.String()
	}
	res := strings.Join(parts, " ")
	return res
}
