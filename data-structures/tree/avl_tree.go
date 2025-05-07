package tree

import "fmt"

type avl_node struct {
	value  int
	height int
	left   *avl_node
	right  *avl_node
}

func new_avl_node(val, height int) *avl_node {
	return &avl_node{
		value:  val,
		height: height,
	}
}

type AVLTree struct {
	size int
	root *avl_node
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

// Check if value 'val' present in the tree
func (tree *AVLTree) Check(val int) bool {
	x := tree.root
	for x != nil {
		if x.value == val {
			return true
		}
		if val < x.value {
			x = x.left
		} else {
			x = x.right
		}
	}
	return false
}

// Minimum value in the tree
func (tree *AVLTree) Minimum() int {
	x := tree.root
	ans := -1
	for x != nil {
		ans = x.value
		x = x.left
	}
	return ans
}

// Maximum value in the tree
func (tree *AVLTree) Maximum() int {
	x := tree.root
	ans := -1
	for x != nil {
		ans = x.value
		x = x.right
	}
	return ans
}

// Minimum value that >= val
func (tree *AVLTree) MinGreater(val int) (int, bool) {
	x := tree.root
	ok := false
	ans := 0
	for x != nil {
		if x.value == val {
			return val, true
		}
		if val < x.value {
			ok = true
			ans = x.value
			x = x.left
		} else {
			x = x.right
		}
	}
	return ans, ok
}

// Maximum value that <= val
func (tree *AVLTree) MaxSmaller(val int) (int, bool) {
	x := tree.root
	ok := false
	ans := 0
	for x != nil {
		if x.value == val {
			return val, true
		}
		if val > x.value {
			ok = true
			ans = x.value
			x = x.right
		} else {
			x = x.left
		}
	}
	return ans, ok
}

// Insert value 'val' to the tree
func (tree *AVLTree) Insert(val int) {

	tree.size++

	if tree.root == nil {
		tree.root = new_avl_node(val, 1)
		return
	}

	tree.root = tree.insert_fixup(tree.root, val)
}

func (tree *AVLTree) Delete(val int) {
	tree.size--
	tree.root = tree.delete_fixup(tree.root, val)
}

// Return size of the tree
func (tree *AVLTree) Size() int {
	return tree.size
}

// Print tree
func (tree *AVLTree) PrintTree() {
	tree.print_node(tree.root, "", true)
}

/*
 */

// <========== - ==========>

func (tree *AVLTree) get_balance_factor(node *avl_node) int {
	if node == nil {
		return 0
	}
	return tree.get_height(node.left) - tree.get_height(node.right)
}

func (tree *AVLTree) get_height(node *avl_node) int {
	if node == nil {
		return 0
	}
	return node.height
}

func (tree *AVLTree) update_height(node *avl_node) {
	h := tree.get_height(node.left)
	r := tree.get_height(node.right)

	if h < r {
		h = r
	}

	node.height = h + 1
}

func (tree *AVLTree) left_rotate(node *avl_node) *avl_node {
	B := node.right
	Y := B.left

	B.left = node
	node.right = Y

	tree.update_height(node)
	tree.update_height(B)

	return B
}

func (tree *AVLTree) right_rotate(node *avl_node) *avl_node {
	A := node.left

	Y := A.right

	A.right = node
	node.left = Y

	tree.update_height(node)
	tree.update_height(A)

	return A
}

func (tree *AVLTree) insert_fixup(node *avl_node, val int) *avl_node {
	if node == nil {
		return new_avl_node(val, 1)
	} else if val < node.value {
		node.left = tree.insert_fixup(node.left, val)
	} else {
		node.right = tree.insert_fixup(node.right, val)
	}

	tree.update_height(node)

	bf := tree.get_balance_factor(node)

	if bf > 1 && val < node.left.value {
		// left heavy
		return tree.right_rotate(node)
	}
	if bf < -1 && val > node.right.value {
		// right heavy
		return tree.left_rotate(node)
	}
	if bf > 1 && val > node.left.value {
		node.left = tree.left_rotate(node.left)
		// left heavy after rotate
		return tree.right_rotate(node)
	}
	if bf < -1 && val < node.right.value {
		node.right = tree.right_rotate(node.right)
		// right heavy after rotate
		return tree.left_rotate(node)
	}

	return node
}

func (tree *AVLTree) subtree_minimum(node *avl_node) *avl_node {
	if node == nil || node.left == nil {
		return node
	}
	return tree.subtree_minimum(node.left)
}

func (tree *AVLTree) delete_fixup(node *avl_node, val int) *avl_node {
	if node == nil {
		return node
	}

	if val < node.value {
		node.left = tree.delete_fixup(node.left, val)
	} else if val > node.value {
		node.right = tree.delete_fixup(node.right, val)
	} else {
		if node.left == nil {
			ret := node.right
			node = nil
			return ret
		}
		if node.right == nil {
			ret := node.left
			node = nil
			return ret
		}
		temp := tree.subtree_minimum(node.right)
		node.value = temp.value
		node.right = tree.delete_fixup(node.right, temp.value)
	}

	tree.update_height(node)

	bf := tree.get_balance_factor(node)

	if bf > 1 && tree.get_balance_factor(node.left) >= 0 {
		// left heavy
		return tree.right_rotate(node)
	}
	if bf < -1 && tree.get_balance_factor(node.right) <= 0 {
		// right heavy
		return tree.left_rotate(node)
	}
	if bf > 1 && tree.get_balance_factor(node.left) < 0 {
		node.left = tree.left_rotate(node.left)
		// left heavy after rotate
		return tree.right_rotate(node)
	}
	if bf < -1 && tree.get_balance_factor(node.right) > 0 {
		node.right = tree.right_rotate(node.right)
		// right heavy after rotate
		return tree.left_rotate(node)
	}

	return node
}

func (tree *AVLTree) print_node(node *avl_node, indent string, isRight bool) {
	if node == nil {
		return
	}
	tree.print_node(node.right, fmt.Sprintf("%s     ", indent), true)
	if isRight {
		fmt.Printf("%s ┌── %d\n", indent, node.value)
	} else {
		fmt.Printf("%s └── %d\n", indent, node.value)
	}
	tree.print_node(node.left, fmt.Sprintf("%s     ", indent), false)
}
