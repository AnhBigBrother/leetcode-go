package tree

import (
	"math/rand"
	"slices"
	"testing"
)

const TEST_ARRAY_LENGTH = 1000

func TestRedBlackTree(t *testing.T) {
	test_case := []int{}
	for range TEST_ARRAY_LENGTH {
		x := rand.Int()
		test_case = append(test_case, x)
	}

	expected := []int{}
	expected = append(expected, test_case...)
	slices.Sort(expected)

	tree := NewRedBlackTree()
	for _, x := range test_case {
		tree.Insert(x)
		if !tree.Check(x) {
			t.Errorf("after insertion, value %d not exist in the tree", x)
		}
	}

	for i := range TEST_ARRAY_LENGTH - 1 {
		if expected[i] == expected[i+1] || expected[i] + 1 == expected[i+1]{
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

	n := TEST_ARRAY_LENGTH
	for tree.Size() > 0 {
		if tree.Size() != n {
			t.Errorf("tree.Size() must be %d, but got %d", n, tree.Size())
		}

		x := tree.Minimum()
		idx := TEST_ARRAY_LENGTH - n

		if x != expected[idx] {
			t.Errorf("%dth minimum element must be %d, but got %d", idx, expected[idx], x)
		}

		tree.Delete(x)
		n--
	}

}
