package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	children := make(map[int]*Node)
	parents := make(map[int]*Node)
	for _, r := range records {
		parent, ok := parents[r.Parent]
		if !ok {
			parents[r.Parent] = &Node{ID: r.Parent}
			parent = parents[r.Parent]
		}
		if r.ID == 0 {
			if r.Parent != 0 {
				return nil, errors.New("root node has parent")
			}
			children[r.ID] = parent
		} else {
			current, ok := parents[r.ID]
			if !ok {
				current = &Node{ID: r.ID}
			}
			parent.Children = append(parent.Children, current)
			sort.Slice(parent.Children, func(i, j int) bool {
				return parent.Children[i].ID < parent.Children[j].ID
			})
		}
	}
	root, ok := children[0]
	if !ok {
		return nil, errors.New("no root node")
	}
	if numberOfNodes(root) != len(records) {
		return nil, errors.New("invalid tree")
	}
	return root, nil
}

func numberOfNodes(node *Node, alreadySeen ...int) int {
	for _, n := range alreadySeen {
		if n == node.ID {
			return 0
		}
	}
	if node.Children == nil {
		return 1
	}
	res := 1
	for _, c := range node.Children {
		res += numberOfNodes(c, append(alreadySeen, node.ID)...)
	}
	return res
}
