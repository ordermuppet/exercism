package robotname

import (
	"errors"
)

var rootNode RuneNode

// Robot is a robot with a unique name.
type Robot struct {
	name string
}

// Name returns a unique name for all Robots created.
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if rootNode.Exhausted {
			return "", errors.New("Possible names exhausted")
		}
		r.name = string(rootNode.RandomSequence())
	}
	return r.name, nil
}

// Reset resets this robot's name.
func (r *Robot) Reset() {
	r.name = ""
}

func init() {
	rootNode = RuneNode{Depth: 0, Children: make(map[rune]*RuneNode)}
}
