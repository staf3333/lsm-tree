package memtable

type node struct {
	key string
	value string
	color color
	left *node
	right *node
	parent *node
}

type RBTree struct {
	root *node
	sentinel *node
}

type color int
const (
	red = iota
	black
)

func NewRBTree() *RBTree {
	sentinel := &node{color: black}
	return &RBTree{
		root: sentinel,
		sentinel: sentinel,
	}
}


func (t *RBTree) leftRotate(node *node) {
	pivot := node.right
	parent := node.parent
	leftSubChild := pivot.left

	node.right = leftSubChild
	leftSubChild.parent = node

	pivot.parent = parent
	pivot.left = node
	if t.root == node {
		t.root = pivot
	} else if parent.left == node {
		parent.left = pivot 
	} else {
		parent.right = pivot
	}

	node.parent = pivot
}

func (t *RBTree) rightRotate(node *node) {
	pivot := node.left
	parent := node.parent
	rightSubChild := pivot.right

	node.left = rightSubChild
	rightSubChild.parent = node

	pivot.parent = parent
	pivot.right = node
	if t.root == node {
		t.root = pivot
	} else if parent.left == node {
		parent.left = pivot 
	} else {
		parent.right = pivot
	}

	node.parent = pivot
}
