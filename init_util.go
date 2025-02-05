package main

import "github.com/AnhBigBrother/leetcode-go/utils"

// <----------------utilities---------------->

// heap, priority queue utils
type MinHeap = utils.MinHeap
type MaxHeap = utils.MaxHeap
type PQItem = utils.PQItem
type PriorityQueue = utils.PriorityQueue

// linked-list, tree utils
type ListNode = utils.ListNode
type TreeNode = utils.TreeNode

var MakeList = utils.MakeList
var ListToArr = utils.ListToArr
var MakeTree = utils.MakeTree
var SpreadTree = utils.SpreadTree

// disjoint set
type DisjointSet = utils.DisjointSet

var NewDisjointSet = utils.NewDisjointSet

// string search utils
var FindLPS = utils.FindLPS
var KmpSearch = utils.KmpSearch

// other
var FindGCD = utils.FindGCD
var FindLCM = utils.FindLCM
var FindPrimeNums = utils.FindPrimeNums
