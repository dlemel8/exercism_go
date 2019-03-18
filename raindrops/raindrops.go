package raindrops

import (
	"bytes"
	"strconv"
)

const testVersion = 3

func appendValueIfDevided(in, num int, val string, b *bytes.Buffer) {
	if in%num == 0 {
		b.WriteString(val)
	}
}

func Convert(in int) string {
	var res bytes.Buffer

	appendValueIfDevided(in, 3, "Pling", &res)
	appendValueIfDevided(in, 5, "Plang", &res)
	appendValueIfDevided(in, 7, "Plong", &res)

	if res.Len() == 0 {
		res.WriteString(strconv.Itoa(in))
	}
	return res.String()
}
