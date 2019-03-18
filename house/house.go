package house

import "strings"

const testVersion = 1

func who(num int) string {
	return []string{
		"",
		"house that Jack built.",
		"malt",
		"rat",
		"cat",
		"dog",
		"cow with the crumpled horn",
		"maiden all forlorn",
		"man all tattered and torn",
		"priest all shaven and shorn",
		"rooster that crowed in the morn",
		"farmer sowing his corn",
		"horse and the hound and the horn",
	}[num]
}

func verb(num int) string {
	return []string{
		"",
		"",
		"lay in",
		"ate",
		"killed",
		"worried",
		"tossed",
		"milked",
		"kissed",
		"married",
		"woke",
		"kept",
		"belonged to",
	}[num]
}

func suffix(num int) string {
	if num == 1 {
		return ""
	}
	return "that " + verb(num) + " the " + who(num-1) + newLine(num, 3) + suffix(num-1)
}

func newLine(num int, from int) string {
	if num >= from {
		return "\n"
	}
	return ""
}

func Verse(num int) string {
	return "This is the " + who(num) + newLine(num, 2) + suffix(num)
}

func Song() string {
	res := make([]string, 12)
	for i := 1; i <= 12; i++ {
		res[i-1] = Verse(i)
	}
	return strings.Join(res, "\n\n")
}
