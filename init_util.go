package main

import (
	"github.com/AnhBigBrother/leetcode-go/algorithms/common"
	"github.com/AnhBigBrother/leetcode-go/algorithms/searching"
	my_heap "github.com/AnhBigBrother/leetcode-go/data-structures/heap"
	"github.com/AnhBigBrother/leetcode-go/data-structures/set"
	"github.com/AnhBigBrother/leetcode-go/data-structures/tree"
	"github.com/AnhBigBrother/leetcode-go/utils"
)

/*
<========== Data Structures ==========>
*/

// heap + priority queue
type MinHeap = my_heap.MinHeap
type MaxHeap = my_heap.MaxHeap
type PQItem = my_heap.PQItem
type PriorityQueue = my_heap.PriorityQueue

// set
type DisjointSet = set.DisjointSet

var NewDisjointSet = set.NewDisjointSet

// tree
type BalanceBinaryTree = tree.BalanceBinaryTree
type RedBlackTree = tree.RedBlackTree

var NewRedBlackTree = tree.NewRedBlackTree

/*
 */

/*
<========== Algorithms ==========>
*/

// string search
var FindLPS = searching.FindLPS
var KmpSearch = searching.KmpSearch

// others
var FindGCD = common.FindGCD
var FindLCM = common.FindLCM
var FindPrimeNums = common.FindPrimeNums
var FindDivisors = common.FindDivisors
var Abs = common.Abs

/*
 */

/*
<========== Utilities ==========>
*/

// linked-list + tree
type ListNode = utils.ListNode
type TreeNode = utils.TreeNode

var NewLinkedList = utils.NewLinkedList
var LinkedListToArr = utils.LinkedListToArr
var NewTree = utils.NewTree
var SpreadTree = utils.SpreadTree
