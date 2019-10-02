package cryptosquare

import (
	"log"
	"math"
	"regexp"
	"strings"
)

// Encode returns a square-coded version of the input text
func Encode(input string) string {
	normalized := normalize(input)
	c := int(math.Ceil(math.Sqrt(float64(len(normalized)))))
	r := c
	if len(normalized) <= c*(c-1) {
		r--
	}

	coded := ""
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			if c*j+i < len(normalized) {
				coded += string(normalized[c*j+i])
			} else {
				coded += " "
			}
		}
		if i < c-1 {
			coded += " "
		}
	}

	return coded
}

func normalize(input string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	return strings.ToLower(reg.ReplaceAllString(input, ""))
}
