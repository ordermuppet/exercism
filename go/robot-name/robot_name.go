package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var usedNames = map[string]bool{}

// Robot is a robot with a unique name.
type Robot struct {
	name string
}

// Name returns a unique name for all Robots created.
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if len(usedNames) == 10*10*10*26*26 {
		return "", errors.New("Possible names exhausted")
	}

	r.name = randomName()
	for usedNames[r.name] {
		r.name = randomName()
	}
	usedNames[r.name] = true
	return r.name, nil
}

// Reset resets this robot's name.
func (r *Robot) Reset() {
	r.name = ""
}

func randomName() string {
	return fmt.Sprintf(
		"%c%c%03d",
		rand.Intn(26)+'A',
		rand.Intn(26)+'A',
		rand.Intn(1000),
	)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
