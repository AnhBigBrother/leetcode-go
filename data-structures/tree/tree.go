package tree

type BalanceBinaryTree interface {
	Check(val int) bool
	MinGreater(val int) (int, bool)
	MaxSmaller(val int) (int, bool)
	Insert(val int)
	Delete(val int)
	Minimum() int
	Maximum() int
	Size() int
	PrintTree()
}
