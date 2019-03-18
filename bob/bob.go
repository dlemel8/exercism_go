package bob

import "regexp"

const testVersion = 3

var yell = regexp.MustCompile(`^[^a-z]*[A-Z]{2,}[^a-z]*$`)
var question = regexp.MustCompile(`^.*\?\s*$`)
var nothing = regexp.MustCompile(`^\s*$`)

func Hey(in string) string {
	if yell.MatchString(in) {
		return "Whoa, chill out!"
	}
	if question.MatchString(in) {
		return "Sure."
	}
	if nothing.MatchString(in) {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
