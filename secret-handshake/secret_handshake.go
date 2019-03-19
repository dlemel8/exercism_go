package secret

const testVersion = 2

const (
	WINK_FLAG            = 1 << 0
	DOUBLE_BLINK_FLAG    = 1 << 1
	CLOSE_YOUR_EYES_FLAG = 1 << 2
	JUMP_FLAG            = 1 << 3
	REVERSE_FLAG         = 1 << 4
)

func isFlagSet(code uint64, flag uint64) bool {
	return code&flag != 0
}

func reverse(ptr *[]string) []string {
	in := *ptr
	last := len(in) - 1
	for i := 0; i < len(in)/2; i++ {
		in[i], in[last-i] = in[last-i], in[i]
	}
	return in
}

func Handshake(code uint) []string {
	res := make([]string, 0, 4)
	if isFlagSet(uint64(code), WINK_FLAG) {
		res = append(res, "wink")
	}
	if isFlagSet(uint64(code), DOUBLE_BLINK_FLAG) {
		res = append(res, "double blink")
	}
	if isFlagSet(uint64(code), CLOSE_YOUR_EYES_FLAG) {
		res = append(res, "close your eyes")
	}
	if isFlagSet(uint64(code), JUMP_FLAG) {
		res = append(res, "jump")
	}
	if isFlagSet(uint64(code), REVERSE_FLAG) {
		reverse(&res)
	}
	return res
}
