package tree

import "fmt"

// refer to: https://github.com/msambol/dsa/blob/master/trees/red_black_tree.py

type color bool

const BLACK color = true
const RED color = false

type rb_node struct {
	value  int
	color  color
	parent *rb_node
	left   *rb_node
	right  *rb_node
}

func new_rb_node(val int, col color) *rb_node {
	return &rb_node{
		value: val,
		color: col,
	}
}

type RedBlackTree struct {
	root *rb_node
	size int
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

// Check if value 'val' present in the tree
func (tree *RedBlackTree) Check(val int) bool {
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
func (tree *RedBlackTree) Minimum() int {
	x := tree.root
	ans := -1
	for x != nil {
		ans = x.value
		x = x.left
	}
	return ans
}

// Maximum value in the tree
func (tree *RedBlackTree) Maximum() int {
	x := tree.root
	ans := -1
	for x != nil {
		ans = x.value
		x = x.right
	}
	return ans
}

// Insert value 'val' to the tree
func (tree *RedBlackTree) Insert(val int) {
	z := new_rb_node(val, RED)

	x := tree.root
	var y *rb_node = nil

	for x != nil {
		y = x
		if val < x.value {
			x = x.left
		} else {
			x = x.right
		}
	}

	z.parent = y
	if y == nil {
		tree.root = z
	} else if val < y.value {
		y.left = z
	} else {
		y.right = z
	}

	tree.insert_fixup(z)
	tree.size++
}

// Delete value 'val' from tree
func (tree *RedBlackTree) Delete(val int) {
	z := tree.search(val)

	if z == nil {
		return
	}

	y := z
	y_orig_color := y.color
	var x *rb_node

	if z.left == nil {
		// case 1: z.left is nil
		x = z.right
		tree.transplant(z, z.right)
	} else if z.right == nil {
		// case 2: z.right is nil
		x = z.left
		tree.transplant(z, z.left)
	} else {
		// case 3: z.left and z.right are not nil
		y = tree.subtree_minimum(z.right)
		y_orig_color = y.color
		x = y.right

		// connect x to y.parent, replace z by y
		if y.parent == z {
			if x != nil {
				x.parent = y
			}
		} else {
			tree.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}

		// replace z by y
		tree.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}

	if y_orig_color == BLACK && x != nil {
		tree.delete_fixup(x)
	}

	tree.size--
}

// Minimum value that >= val
func (tree *RedBlackTree) MinGreater(val int) (int, bool) {
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
func (tree *RedBlackTree) MaxSmaller(val int) (int, bool) {
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

// Return size of the tree
func (tree *RedBlackTree) Size() int {
	return tree.size
}

// Print tree
func (tree *RedBlackTree) PrintTree() {
	tree.print_node(tree.root, "", true)
}

func (tree *RedBlackTree) left_rotate(x *rb_node) {
	y := x.right
	if y == nil {
		return
	}

	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}

	y.parent = x.parent

	if x.parent == nil {
		tree.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

func (tree *RedBlackTree) right_rotate(x *rb_node) {
	y := x.left
	if y == nil {
		return
	}

	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}

	y.parent = x.parent

	if x.parent == nil {
		tree.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.right = x
	x.parent = y
}

func (tree *RedBlackTree) insert_fixup(z *rb_node) {

	for z.parent != nil && z.parent.color == RED {

		// root node is always BLACK, but z.parent is RED => z has grandparent (z.parent.parent != nil)

		if z.parent == z.parent.parent.left {
			// z.parent is left child of z.grandparent

			// y is uncle of z
			y := z.parent.parent.right

			if y != nil && y.color == RED {
				// case 1 : z 's uncle is RED
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				// case 2: z 's uncle is BLACK and z, z.parent, z.grandparent positions are in triangle
				if z == z.parent.right {
					z = z.parent
					tree.left_rotate(z) // left triangle
				}

				// after rotate, case 2 become case 3

				// case 3: z is RED, z.parent is RED, z.uncle is black and z, z.parent, z.grandparent positions are in line
				z.parent.color = BLACK
				z.parent.parent.color = RED
				tree.right_rotate(z.parent.parent) // to the right line
			}
		} else {
			// z.parent is right child of z.grandparent

			// y is uncle of z
			y := z.parent.parent.left

			if y != nil && y.color == RED {
				// case 1 : z 's uncle is RED
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				// case 2: z 's uncle is BLACK and z, z.parent, z.grandparent positions are in triangle
				if z == z.parent.left {
					z = z.parent
					tree.right_rotate(z) // right triangle
				}

				// after rotate, case 2 become case 3

				// case 3: z is RED, z.parent is RED, z.uncle is black and z, z.parent, z.grandparent positions are in line
				z.parent.color = BLACK
				z.parent.parent.color = RED
				tree.left_rotate(z.parent.parent) // to the left line
			}
		}

		if z == tree.root {
			break
		}

	}

	tree.root.color = BLACK
}

func (tree *RedBlackTree) search(val int) *rb_node {
	x := tree.root
	for x != nil && x.value != val {
		if val < x.value {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

// transplant(u, v) connect v to u.parent and free u
func (tree *RedBlackTree) transplant(u, v *rb_node) {
	if u.parent == nil {
		tree.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func (tree *RedBlackTree) subtree_minimum(x *rb_node) *rb_node {
	for x.left != nil {
		x = x.left
	}
	return x
}

func (tree *RedBlackTree) delete_fixup(x *rb_node) {

	for x != tree.root && x.color == BLACK {

		if x == x.parent.left {

			// w is x's sibling
			w := x.parent.right
			if w == nil {
				return
			}

			// w is RED => x.parent is BLACK
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				tree.left_rotate(x.parent)
				w = x.parent.right
				if w == nil {
					return
				}
			}

			// w.left and w.right are BLACK
			if (w.left == nil || w.left.color == BLACK) && (w.right == nil || w.right.color == BLACK) {

				w.color = RED
				x = x.parent

			} else {

				// w.left is RED and w.right is BLACK
				if w.right == nil || w.right.color == BLACK {
					w.left.color = BLACK
					w.color = RED
					tree.right_rotate(w)
					w = x.parent.right
					if w == nil {
						return
					}
				}

				// w.right is RED
				w.color = x.parent.color
				x.parent.color = BLACK
				w.right.color = BLACK
				tree.left_rotate(x.parent)
				x = tree.root
			}

		} else {

			// w is x's sibling
			w := x.parent.left
			if w == nil {
				return
			}

			// w is RED => x.parent is BLACK
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				tree.right_rotate(x.parent)
				w = x.parent.left
				if w == nil {
					return
				}
			}

			// w.left and w.right are BLACK
			if (w.left == nil || w.left.color == BLACK) && (w.right == nil || w.right.color == BLACK) {

				w.color = RED
				x = x.parent

			} else {

				// w.left is BLACK and w.right is RED
				if w.left == nil || w.left.color == BLACK {
					w.right.color = BLACK
					w.color = RED
					tree.left_rotate(w)
					w = x.parent.left
					if w == nil {
						return
					}
				}

				// w.left is RED
				w.color = x.parent.color
				x.parent.color = BLACK
				w.left.color = BLACK
				tree.right_rotate(x.parent)
				x = tree.root
			}

		}

	}

	x.color = BLACK
}

func (tree *RedBlackTree) print_node(node *rb_node, indent string, isRight bool) {
	if node == nil {
		return
	}
	tree.print_node(node.right, fmt.Sprintf("%s       ", indent), true)
	c := "B"
	if node.color == RED {
		c = "R"
	}
	if isRight {
		fmt.Printf("%s   ┌── [%d, %s]\n", indent, node.value, c)
	} else {
		fmt.Printf("%s   └── [%d, %s]\n", indent, node.value, c)
	}
	tree.print_node(node.left, fmt.Sprintf("%s       ", indent), false)
}
