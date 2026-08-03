package pov

type Tree struct {
	value    string
	parent   *Tree
	children map[string]*Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	t := Tree{
		value: value,
	}

	m := make(map[string]*Tree)
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
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// POV problem-specific functions

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	node, ok := traverse(tr, from)
	if !ok {
		return nil
	}
	return reorient(node, "")
}

// reorient rebuilds the tree rooted at node, treating node's parent chain
// as additional children. excludeChild is the value of the child that the
// recursion just came from, so it isn't duplicated below node.
func reorient(node *Tree, excludeChild string) *Tree {
	children := make([]*Tree, 0, len(node.children)+1)
	for value, c := range node.children {
		if value == excludeChild {
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
	fNode, ok := traverse(tr, from)
	if !ok {
		return nil
	}
	tNode, ok := traverse(tr, to)
	if !ok {
		return nil
	}

	fPath := path2Root(*fNode)
	tPath := path2Root(*tNode)

	fi, ti := len(fPath)-1, len(tPath)-1

	for fi >= 0 && ti >= 0 && fPath[fi].value == tPath[ti].value {
		fi--
		ti--
	}

	path := make([]string, 0, 1)
	for i := 0; i <= fi; i++ {
		path = append(path, fPath[i].value)
	}
	if fi+1 < len(fPath) {
		path = append(path, fPath[fi+1].value)
	} else if ti+1 < len(tPath) {
		path = append(path, tPath[ti+1].value)
	}
	for i := ti; i >= 0; i-- {
		path = append(path, tPath[i].value)
	}
	return path
}

func path2Root(tr Tree) []Tree {
	path := make([]Tree, 0, 1)
	path = append(path, tr)
	for tr.parent != nil {
		path = append(path, *tr.parent)
		tr = *tr.parent
	}
	return path
}

func traverse(tr *Tree, from string) (*Tree, bool) {
	if tr.value == from {
		return tr, true
	}
	if tr.parent != nil && tr.parent.value == from {
		return tr.parent, true
	}
	if node, ok := tr.children[from]; ok {
		return node, true
	}
	for _, c := range tr.children {
		if node, ok := traverse(c, from); ok {
			return node, true
		}
	}
	return nil, false
}
