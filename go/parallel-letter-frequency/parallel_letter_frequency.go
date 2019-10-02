package letter

// ConcurrentFrequency counts the occurrences of a letter in each of the provided texts.
func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan FreqMap, 10)
	for _, text := range texts {
		go func(text string) {
			c <- Frequency(text)
		}(text)
	}
	frequencies := FreqMap{}
	for range texts {
		counts := <-c
		for letter, count := range counts {
			frequencies[letter] += count
		}
	}

	return frequencies
}
