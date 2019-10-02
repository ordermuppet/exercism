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
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	nodes := make(map[int]*Node)
	root, records := records[0], records[1:]

	if root.ID != 0 {
		return &Node{}, errors.New("No root node present")
	}
	if root.Parent != 0 {
		return &Node{}, errors.New("Root node cannot have parent")
	}
	nodes[0] = &Node{ID: 0}

	for i, record := range records {
		if i != record.ID-1 {
			return &Node{}, errors.New("Non-continuous record")
		}
		if record.Parent >= record.ID {
			return &Node{}, errors.New("Cycle detected")
		}

		nodes[record.ID] = &Node{ID: record.ID}

		if parent, ok := nodes[record.Parent]; ok {
			parent.Children = append(parent.Children, nodes[record.ID])
		} else {
			return &Node{}, errors.New("Cannot find parent")
		}
	}

	return nodes[0], nil
}
