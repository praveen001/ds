package binarytree

// BinaryTree represents a binary tree
type BinaryTree struct {
	root *node
	size int
}

// Node represents a node in a binary tree
type node struct {
	value int
	left  *node
	right *node
}

// New creates a new instance of binary tree and returns it
func New() *BinaryTree {
	return &BinaryTree{}
}

// NewNode returns a new binary tree node with given value
func newNode(value int) *node {
	return &node{value, nil, nil}
}

// Insert given values into the tree
func (bt *BinaryTree) Insert(value int) {
	bt.size++

	if bt.root == nil {
		bt.root = newNode(value)
		return
	}

	node := bt.root
	for {
		if node.value < value {
			if node.right == nil {
				node.right = newNode(value)
				break
			} else {
				node = node.right
			}
		} else {
			if node.left == nil {
				node.left = newNode(value)
				break
			} else {
				node = node.left
			}
		}
	}
}

// Contains returns true if the given value exists in the tree, otherwise false
func (bt *BinaryTree) Contains(value int) bool {
	if bt.Size() == 0 {
		return false
	}

	node := bt.root
	for {
		if node.value < value {
			if node.right == nil {
				return false
			}
			node = node.right
		} else if node.value > value {
			if node.left == nil {
				return false
			}
			node = node.left
		} else {
			return true
		}
	}
}

// Min returns the minimum value from the tree
func (bt *BinaryTree) Min() int {
	if bt.Size() == 0 {
		return -1
	}

	node := bt.root
	for {
		if node.left == nil {
			return node.value
		}
		node = node.left
	}
}

// Max returns the maximum value from the tree
func (bt *BinaryTree) Max() int {
	if bt.Size() == 0 {
		return -1
	}

	node := bt.root
	for {
		if node.right == nil {
			return node.value
		}
		node = node.right
	}
}

// Size returns the total number of values in the tree
func (bt *BinaryTree) Size() int {
	return bt.size
}

// InOrder ..
func (bt *BinaryTree) InOrder() []int {
	arr := make([]int, bt.Size())
	index := 0

	if bt.Size() != 0 {
		bt.root.inOrder(arr, &index)
	}

	return arr
}

func (n *node) inOrder(arr []int, index *int) {
	if n.left != nil {
		n.left.inOrder(arr, index)
	}

	arr[*index] = n.value
	*index++

	if n.right != nil {
		n.right.inOrder(arr, index)
	}
}
