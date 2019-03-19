package twelve

import (
	"bytes"
	"fmt"
	"strings"
)

const testVersion = 1

func ordinal(day int) string {
	options := []string{
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
		"sixth",
		"seventh",
		"eighth",
		"ninth",
		"tenth",
		"eleventh",
		"twelfth",
	}
	return options[day-1]
}

func gifts(day int) string {
	options := []string{
		"a Partridge",
		"two Turtle Doves",
		"three French Hens",
		"four Calling Birds",
		"five Gold Rings",
		"six Geese-a-Laying",
		"seven Swans-a-Swimming",
		"eight Maids-a-Milking",
		"nine Ladies Dancing",
		"ten Lords-a-Leaping",
		"eleven Pipers Piping",
		"twelve Drummers Drumming",
	}

	selected := options[:day]
	var res bytes.Buffer
	for i := len(selected) - 1; i >= 0; i-- {
		var format string
		if i == 0 {
			format = "%s"
		} else if i == 1 {
			format = "%s, and "
		} else {
			format = "%s, "
		}
		res.WriteString(fmt.Sprintf(format, selected[i]))
	}

	return res.String()
}

func Verse(day int) string {
	format := "On the %s day of Christmas my true love gave to me, %s in a Pear Tree."
	return fmt.Sprintf(format, ordinal(day), gifts(day))

}

func Song() string {
	verses := make([]string, 12)
	for i := 0; i < 12; i++ {
		verses[i] = Verse(i + 1)
	}
	return strings.Join(verses, "\n") + "\n"
}
