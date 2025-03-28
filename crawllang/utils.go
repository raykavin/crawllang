package crawllang

import "regexp"

// Checks if character is a letter
func isLetter(ch byte) bool {
	match, _ := regexp.MatchString(`[a-zA-Z]`, string(ch))
	return match
}

// Checks if character is a digit
func isDigit(ch byte) bool {
	match, _ := regexp.MatchString(`[0-9]`, string(ch))
	return match
}
