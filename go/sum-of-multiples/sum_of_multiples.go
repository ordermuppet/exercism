package summultiples

// SumMultiples returns the sum of all unique multiples of the divisors below the limit
func SumMultiples(limit int, divisors ...int) int {
	uniques := make(map[int]struct{})
	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}
		for i := divisor; i < limit; i = i + divisor {
			_, ok := uniques[i]
			if !ok {
				uniques[i] = struct{}{}
			}
		}
	}

	sum := 0
	for k := range uniques {
		sum += k
	}
	return sum
}
