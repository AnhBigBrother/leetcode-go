package cache

import (
	"testing"
)

// more tests: https://leetcode.com/problems/lru-cache/description/

func TestLRUcache(t *testing.T) {

	var test_comman = []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	var test_input = [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}
	var expected_result = []int{-1, -1, -1, 1, -1, -1, -1, -1, 3, 4}

	n := len(test_comman)

	cache := NewLRUCache(test_input[0][0])

	for i := 1; i < n; i++ {

		if test_comman[i] == "put" {

			cache.Put(test_input[i][0], test_input[i][1])

		} else if test_comman[i] == "get" {

			val := cache.Get(test_input[i][0])

			if val != expected_result[i] {
				t.Errorf("cache.Get falied, expected %d, but got %d", expected_result[i], val)
			}

		}
	}
}
