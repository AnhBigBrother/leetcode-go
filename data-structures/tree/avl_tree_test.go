package tree

import (
	"math/rand"
	"slices"
	"testing"
)

func TestAVLTree(t *testing.T) {

	const TREE_SIZE = 1000

	test_case := []int{}
	for range TREE_SIZE {
		x := rand.Int()
		test_case = append(test_case, x)
	}

	expected := []int{}
	expected = append(expected, test_case...)
	slices.Sort(expected)

	tree := NewAVLTree()
	for _, x := range test_case {
		tree.Insert(x)
		if !tree.Check(x) {
			t.Errorf("after insertion, value %d not exist in the tree", x)
		}
	}

	for i := range TREE_SIZE - 1 {
		if expected[i] == expected[i+1] || expected[i]+1 == expected[i+1] {
			continue
		}

		x := expected[i]/2 + expected[i+1]/2
		y, ok1 := tree.MaxSmaller(x)
		z, ok2 := tree.MinGreater(x)

		if !(ok1 && ok2) {
			t.Errorf("cannot find MinGreater and MaxSmaller of valid value: %d", x)
		}

		if y != expected[i] {
			t.Errorf("wrong MinGreater of: %d, expected %d, but got %d", x, expected[i], y)
		}

		if z != expected[i+1] {
			t.Errorf("wrong MaxSmaller of: %d, expected %d, but got %d", x, expected[i+1], z)
		}
	}

	n := TREE_SIZE
	for tree.Size() > 0 {
		if tree.Size() != n {
			t.Errorf("tree.Size() must be %d, but got %d", n, tree.Size())
		}

		x := tree.Minimum()
		idx := TREE_SIZE - n

		if x != expected[idx] {
			t.Errorf("%dth minimum element must be %d, but got %d", idx, expected[idx], x)
		}

		tree.Delete(x)
		n--
	}

}

func BenchmarkAVLTreeInsert100000(b *testing.B) {

	b.StopTimer()

	const TREE_SIZE = 100000

	test_case := []int{}
	for range TREE_SIZE {
		x := rand.Int()
		test_case = append(test_case, x)
	}

	tree := NewAVLTree()

	b.StartTimer()

	for _, x := range test_case {
		tree.Insert(x)
	}

	b.StopTimer()
}

func BenchmarkAVLTreeDelete100000(b *testing.B) {

	b.StopTimer()

	const TREE_SIZE = 100000

	test_case := []int{}
	for range TREE_SIZE {
		x := rand.Int()
		test_case = append(test_case, x)
	}

	tree := NewAVLTree()
	for _, x := range test_case {
		tree.Insert(x)
	}

	b.StartTimer()

	for i := TREE_SIZE - 1; i >= 0; i-- {
		x := test_case[i]
		tree.Delete(x)
	}

	b.StopTimer()
}

func BenchmarkAVLTreeQuery100000(b *testing.B) {

	b.StopTimer()

	const TREE_SIZE = 100000

	test_case := []int{}
	for range TREE_SIZE {
		x := rand.Int()
		test_case = append(test_case, x)
	}

	tree := NewAVLTree()
	for _, x := range test_case {
		tree.Insert(x)

	}

	b.StartTimer()

	for i := TREE_SIZE - 1; i >= 0; i-- {
		x := test_case[i]
		tree.Check(x)
	}

	b.StopTimer()
}
