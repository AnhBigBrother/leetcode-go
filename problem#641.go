package main

type listnode struct {
	value int
	front *listnode
	back  *listnode
}

type MyCircularDeque struct {
	head       *listnode
	tail       *listnode
	length     int
	max_length int
}

func MyCircularDequeConstructor(k int) MyCircularDeque {
	return MyCircularDeque{max_length: k}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.length == this.max_length {
		return false
	}
	node := &listnode{value: value}
	if this.head == nil {
		this.tail = node
	} else if this.length == 1 {
		node.back = this.tail
		this.tail.front = node
	} else {
		node.back = this.head
		this.head.front = node
	}
	this.head = node
	this.length++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.length == this.max_length {
		return false
	}
	node := &listnode{value: value}
	if this.tail == nil {
		this.head = node
	} else if this.length == 1 {
		node.front = this.head
		this.head.back = node
	} else {
		node.front = this.tail
		this.tail.back = node
	}
	this.tail = node
	this.length++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.head == nil {
		return false
	} else if this.length == 1 {
		this.head = nil
		this.tail = nil
	} else {
		this.head = this.head.back
	}
	this.length--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.tail == nil {
		return false
	} else if this.length == 1 {
		this.head = nil
		this.tail = nil
	} else {
		this.tail = this.tail.front
	}
	this.length--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.head == nil {
		return -1
	}
	return this.head.value
}

func (this *MyCircularDeque) GetRear() int {
	if this.tail == nil {
		return -1
	}
	return this.tail.value
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.length == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.length == this.max_length
}
