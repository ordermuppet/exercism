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

	possibleRunes := make([]rune, 0)
	if c.Depth == 0 || c.Depth == 1 {
		for char := 'A'; char <= 'Z'; char++ {
			if node, ok := c.Children[char]; ok {
				if !node.Exhausted {
					possibleRunes = append(possibleRunes, char)
				}
			} else {
				possibleRunes = append(possibleRunes, char)
			}
		}
	} else {
		for char := '0'; char <= '9'; char++ {
			if node, ok := c.Children[char]; ok {
				if !node.Exhausted {
					possibleRunes = append(possibleRunes, char)
				}
			} else {
				possibleRunes = append(possibleRunes, char)
			}
		}
	}

	// If there's only one possible rune left, this node is now exhausted
	var r rune
	// fmt.Println(possibleRunes)
	if len(possibleRunes) == 1 {
		// c.Exhausted = true
		r = possibleRunes[0]
	} else {
		// fmt.Println(len(possibleRunes))
		r = possibleRunes[rand.Intn(len(possibleRunes))]
	}
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
	} else {
		return append([]byte{byte(c.Char)}, remainder...)
	}
}

// Name returns a unique name for all Robots created.
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if rootNode.Exhausted {
			// fmt.Println(rootNode)
			return "", errors.New("Possible names exhausted")
		}
		r.name = string(rootNode.sequence())
	}
	return r.name, nil
}

// Reset resets this robot's name.
func (r *Robot) Reset() {
	r.name = ""
}

func init() {
	rootNode = CharacterNode{Depth: 0, Children: make(map[rune]*CharacterNode)}
}
