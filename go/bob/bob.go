package bob

import "strings"

// Remark represents something said to Bob the teenager.
type Remark string

// Hey returns the response of Bob the teenager.
func Hey(remark string) string {
	r := Remark(remark)
	switch {
	case r.isQuestion() && r.isYelling():
		return "Calm down, I know what I'm doing!"
	case r.isQuestion():
		return "Sure."
	case r.isYelling():
		return "Whoa, chill out!"
	case r.isEmpty():
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}

func (r Remark) isYelling() bool {
	hasLetter := false
	for _, c := range string(r) {
		if c >= 'a' && c <= 'z' {
			return false
		}
		if c >= 'A' && c <= 'Z' {
			hasLetter = true
		}
	}
	return string(r) == strings.ToUpper(string(r)) && hasLetter
}

func (r Remark) isQuestion() bool {
	return strings.HasSuffix(strings.TrimSpace(string(r)), "?")
}

func (r Remark) isEmpty() bool {
	return len(strings.TrimSpace(string(r))) == 0
}
