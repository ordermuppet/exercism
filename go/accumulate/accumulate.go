package accumulate

// Accumulate returns an array that holds the result of applying f to each element in input
func Accumulate(input []string, f func(string) string) []string {
	output := make([]string, len(input))
	for i, a := range input {
		output[i] = f(a)
	}
	return output
}
