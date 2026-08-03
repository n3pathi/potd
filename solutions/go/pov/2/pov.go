package pov

import "strings"

type Tree struct {
	value    string
	parent   *Tree
	children map[string]*Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	t := Tree{value: value}

	m := make(map[string]*Tree, len(children))
	for _, c := range children {
		m[c.value] = c
		c.parent = &t
	}
	t.children = m

	return &t
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.value
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	children := make([]*Tree, 0, len(tr.children))
	for _, v := range tr.children {
		children = append(children, v)
	}
	return children
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	children := tr.Children()
	if len(children) == 0 {
		return tr.Value()
	}
	parts := make([]string, 0, len(children)+1)
	parts = append(parts, tr.Value())
	for _, ch := range children {
		parts = append(parts, ch.String())
	}
	return "(" + strings.Join(parts, " ") + ")"
}

// POV problem-specific functions

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	node, ok := find(tr, from)
	if !ok {
		return nil
	}
	return reorient(node, "")
}

// reorient rebuilds the tree rooted at node, treating node's parent chain
// as additional children. cameFrom is the value of the child that the
// recursion just came from, so it isn't duplicated below node.
func reorient(node *Tree, cameFrom string) *Tree {
	children := make([]*Tree, 0, len(node.children)+1)
	for value, c := range node.children {
		if value == cameFrom {
			continue
		}
		children = append(children, c)
	}
	if node.parent != nil {
		children = append(children, reorient(node.parent, node.value))
	}
	return New(node.value, children...)
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	if from == "" || to == "" {
		return nil
	}
	fNode, ok := find(tr, from)
	if !ok {
		return nil
	}
	tNode, ok := find(tr, to)
	if !ok {
		return nil
	}

	// ancestry lists node values from the node itself up to the root, so
	// the last common values (scanning from the end) mark the lowest
	// common ancestor of fNode and tNode.
	fPath := ancestry(fNode)
	tPath := ancestry(tNode)

	fi, ti := len(fPath)-1, len(tPath)-1
	for fi >= 0 && ti >= 0 && fPath[fi] == tPath[ti] {
		fi--
		ti--
	}

	path := append([]string{}, fPath[:fi+2]...) // from-node ... lowest common ancestor
	for i := ti; i >= 0; i-- {                  // ... down to to-node
		path = append(path, tPath[i])
	}
	return path
}

// ancestry returns the values from node up to the root, inclusive.
func ancestry(node *Tree) []string {
	values := []string{node.value}
	for node.parent != nil {
		node = node.parent
		values = append(values, node.value)
	}
	return values
}

// find searches the whole tree containing tr (walking up to the real root
// first, in case tr isn't it) for a node with the given value.
func find(tr *Tree, value string) (*Tree, bool) {
	root := tr
	for root.parent != nil {
		root = root.parent
	}
	return search(root, value)
}

func search(tr *Tree, value string) (*Tree, bool) {
	if tr.value == value {
		return tr, true
	}
	for _, c := range tr.children {
		if node, ok := search(c, value); ok {
			return node, true
		}
	}
	return nil, false
}
