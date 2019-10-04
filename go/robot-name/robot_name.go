package robotname

import (
	"errors"
	"math/rand"
)

var rootNode CharacterNode

// Robot is a robot with a unique name.
type Robot struct {
	name string
}

// CharacterNode is a thing
type CharacterNode struct {
	Char      rune
	Children  map[rune]*CharacterNode
	Exhausted bool
	Depth     int
}

func (c *CharacterNode) sequence() []byte {
	if c.Depth == 5 {
		c.Exhausted = true
		return []byte{byte(c.Char)}
	}

	possibleRunes := c.possibleRunes()
	r := possibleRunes[rand.Intn(len(possibleRunes))]
	var selectedNode *CharacterNode
	if node, ok := c.Children[r]; ok {
		selectedNode = node
	} else {
		selectedNode = &CharacterNode{Depth: c.Depth + 1, Char: r, Children: make(map[rune]*CharacterNode)}
	}

	c.Children[r] = selectedNode
	remainder := selectedNode.sequence()
	if len(possibleRunes) == 1 && selectedNode.Exhausted {
		c.Exhausted = true
	}

	if c.Depth == 0 {
		return remainder
	}
	return append([]byte{byte(c.Char)}, remainder...)
}

// Name returns a unique name for all Robots created.
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if rootNode.Exhausted {
			return "", errors.New("Possible names exhausted")
		}
		r.name = string(rootNode.sequence())
	}
	return r.name, nil
}

func (c CharacterNode) possibleRunes() []rune {
	runes := make([]rune, 0)

	var min, max rune
	if c.Depth == 0 || c.Depth == 1 {
		min, max = 'A', 'Z'
	} else {
		min, max = '0', '9'
	}

	for r := min; r <= max; r++ {
		if node, ok := c.Children[r]; !ok || !node.Exhausted {
			runes = append(runes, r)
		}
	}

	return runes
}

// Reset resets this robot's name.
func (r *Robot) Reset() {
	r.name = ""
}

func init() {
	rootNode = CharacterNode{Depth: 0, Children: make(map[rune]*CharacterNode)}
}
