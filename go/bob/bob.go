package bob

import "regexp"

var reQuestion = regexp.MustCompile(`\?\s*$`)
var reYell = regexp.MustCompile(`^[^a-z]*[A-Z]+[^a-z?]*$`)
var reYellQuestion = regexp.MustCompile(`^[^a-z]*[A-Z]+[^a-z]*\?\s*$`)
var reSilence = regexp.MustCompile(`^\s*$`)

func Hey(remark string) string {

	switch {
	case reYellQuestion.MatchString(remark):
		return "Calm down, I know what I'm doing!"
	case reQuestion.MatchString(remark):
		return "Sure."
	case reYell.MatchString(remark):
		return "Whoa, chill out!"
	case reSilence.MatchString(remark):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
