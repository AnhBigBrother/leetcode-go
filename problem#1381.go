package main

type CustomStack struct {
	stack    []int
	inc      []int
	max_size int
}

func CustomStackConstructor(maxSize int) CustomStack {
	return CustomStack{
		max_size: maxSize,
	}
}

func (this *CustomStack) Push(x int) {
	if len(this.stack) == this.max_size {
		return
	}
	this.stack = append(this.stack, x)
	this.inc = append(this.inc, 0)
}

func (this *CustomStack) Pop() int {
	res := 0
	i := len(this.stack) - 1
	if i < 0 {
		return -1
	}
	if i > 0 {
		this.inc[i-1] += this.inc[i]
	}
	res += this.inc[i] + this.stack[i]
	this.stack = this.stack[:i]
	this.inc = this.inc[:i]
	return res
}

func (this *CustomStack) Increment(k int, val int) {
	i := min(k, len(this.stack)) - 1
	if i >= 0 {
		this.inc[i] += val
	}
}
