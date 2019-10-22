package robotname

import (
	"math/rand"
)

// RuneNode represents a node in a tree of runes that make up a robot name.
// This data structure allows us to efficiently check for already-used robot names
// without needing to generate all possible names ahead of time.
type RuneNode struct {
	Rune      rune
	Children  map[rune]*RuneNode
	Exhausted bool
	Depth     int
}

// RandomSequence returns a random sequence of runes starting from the current node.
func (rn *RuneNode) RandomSequence() []rune {
	if rn.Depth == 5 {
		// 5 is the max depth of the sequence, so is exhausted by default (has no possible children)
		rn.Exhausted = true
		return []rune{rn.Rune}
	}

	possibleRunes := rn.possibleRunes()
	r := possibleRunes[rand.Intn(len(possibleRunes))]
	var selectedNode *RuneNode
	if node, ok := rn.Children[r]; ok {
		selectedNode = node
	} else {
		selectedNode = &RuneNode{Depth: rn.Depth+1, Rune: r, Children: make(map[rune]*RuneNode)}
		rn.Children[r] = selectedNode
	}
	
	remainder := selectedNode.RandomSequence()

	// if there was only one possible rune and its node is now exhausted, then the current
	// node has now become exhausted
	if len(possibleRunes) == 1 && selectedNode.Exhausted {
		rn.Exhausted = true
	}

	if rn.Depth == 0 {
		return remainder
	}
	return append([]rune{rn.Rune}, remainder...)
}

func (rn RuneNode) possibleRunes() []rune {
	runes := make([]rune, 0)

	var min, max rune
	if rn.Depth == 0 || rn.Depth == 1 {
		min, max = 'A', 'Z'
	} else {
		min, max = '0', '9'
	}

	for r := min; r <= max; r++ {
		if node, ok := rn.Children[r]; !ok || !node.Exhausted {
			runes = append(runes, r)
		}
	}

	return runes
}