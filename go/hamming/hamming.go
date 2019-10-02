package hamming

import "errors"

// Distance returns the number of differences between the two strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strands must be the same length")
	}

	count := 0
	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
