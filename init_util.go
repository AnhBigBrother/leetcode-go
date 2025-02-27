package main

import "github.com/AnhBigBrother/leetcode-go/utils"

// <----------------utilities---------------->

// heap + priority queue
type MinHeap = utils.MinHeap
type MaxHeap = utils.MaxHeap
type PQItem = utils.PQItem
type PriorityQueue = utils.PriorityQueue

// linked-list + tree
type ListNode = utils.ListNode
type TreeNode = utils.TreeNode

var NewLinkedList = utils.NewLinkedList
var LinkedListToArr = utils.LinkedListToArr
var NewTree = utils.NewTree
var SpreadTree = utils.SpreadTree

// disjoint set
type DisjointSet = utils.DisjointSet

var NewDisjointSet = utils.NewDisjointSet

// string search
var FindLPS = utils.FindLPS
var KmpSearch = utils.KmpSearch

// others
var FindGCD = utils.FindGCD
var FindLCM = utils.FindLCM
var FindPrimeNums = utils.FindPrimeNums
var FindDivisors = utils.FindDivisors
var Abs = utils.Abs
