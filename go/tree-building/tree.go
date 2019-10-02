package tree

import (
	"errors"
	"sort"
)

// Record represents a persisted post in a hypothetical web-forum.
type Record struct {
	ID     int
	Parent int
}

// Node represents a web forum post.
type Node struct {
	ID       int
	Children []*Node
}

// Build creates a tree structure from a flat array of records
func Build(records []Record) (*Node, error) {
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	nodes := make(map[int]*Node)
	for i, record := range records {
		if i != record.ID {
			return &Node{}, errors.New("Non-continuous record")
		}
		if i == 0 && record.Parent != 0 {
			return &Node{}, errors.New("Root node cannot have parent")
		}
		if record.Parent >= i && i != 0 {
			return &Node{}, errors.New("Cycle detected")
		}
		nodes[i] = &Node{ID: i}
		if i != 0 {
			nodes[record.Parent].Children = append(nodes[record.Parent].Children, nodes[i])
		}
	}
	return nodes[0], nil
}
