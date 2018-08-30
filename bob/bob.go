package bob

import (
	"strings"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	switch {
	case isYell(remark) && isQuestion(remark):
		return "Calm down, I know what I'm doing!"
	case isYell(remark):
		return "Whoa, chill out!"
	case isQuestion(remark):
		return "Sure."
	case remark == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}

}

func isYell(phrase string) bool {
	return phrase == strings.ToUpper(phrase) && phrase != strings.ToLower(phrase)
}

func isQuestion(phrase string) bool {
	return strings.HasSuffix(phrase, "?")
}
