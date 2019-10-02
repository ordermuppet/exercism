package reverse

import "strings"

// Reverse returns the string backwards
func Reverse(input string) string {
	r := []rune(input)
	var reversedBuilder strings.Builder
	for i := len(r) - 1; i >= 0; i-- {
		reversedBuilder.WriteRune(r[i])
	}
	return reversedBuilder.String()
}
