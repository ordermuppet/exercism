package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

var usedNames map[string]struct{}

// Robot is a robot with a unique name.
type Robot struct {
	name string
}

// Name returns a unique name for all Robots created.
func (r *Robot) Name() (string, error) {
	if len(usedNames) == 10*10*10*26*26 { // 10 possible digits, 26 possible letters
		return "", errors.New("Possible names exhausted")
	}
	if r.name == "" {
		for {
			r.name = randomName()
			if _, ok := usedNames[r.name]; !ok {
				usedNames[r.name] = struct{}{}
				break
			}
		}
	}
	fmt.Println(r.name)
	return r.name, nil
}

// Reset resets this robot's name.
func (r *Robot) Reset() {
	r.name = ""
}

func randomName() string {
	return string([]byte{
		randomLetter(),
		randomLetter(),
		randomNumber(),
		randomNumber(),
		randomNumber(),
	})
}

func randomLetter() byte {
	return byte(65 + rand.Intn(25)) // 65 = A, 90 = Z
}

func randomNumber() byte {
	return byte(48 + rand.Intn(9)) // 48 = string '0', 57 = string '9'
}

func init() {
	usedNames = make(map[string]struct{})
}
