package datastructures

type Leaf struct {
	Value float64
	//otherwise the compiler have no method to determine the size of the stack
	Left   *Leaf
	Right  *Leaf
	Parent *Leaf
}

type Tree struct {
	Root *Leaf
}

func (t Tree) Insert(val float64) {
	t.insertHelper(t.Root, val)
}

func (t Tree) insertHelper(node *Leaf, val float64) *Leaf {
	if t.Root == nil {
		t.Root = &Leaf{Value: val}
		return t.Root
	}
	if node == nil {
		return &Leaf{Value: val}
	}
	if val <= node.Value {
		node.Left = t.insertHelper(node.Left, val)
	}
	if val > node.Value {
		node.Right = t.insertHelper(node.Right, val)
	}
	return node
}