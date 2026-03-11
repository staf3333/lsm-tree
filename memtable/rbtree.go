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
