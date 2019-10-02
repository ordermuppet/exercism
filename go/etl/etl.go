package etl

import "strings"

// Transform converts a map of scores to letters with that score, to
// a map of letters to scores
func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)
	for score, letters := range input {
		for _, letter := range letters {
			output[strings.ToLower(letter)] = score
		}
	}
	return output
}
