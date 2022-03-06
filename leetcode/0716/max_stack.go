package leetcode

type MaxStack struct {
	stack    []int
	maxStack []int
}

func Constructor() MaxStack {
	return MaxStack{
		stack:    []int{},
		maxStack: []int{},
	}
}

func (this *MaxStack) Push(x int) {
	if len(this.maxStack) == 0 || this.maxStack[len(this.maxStack)-1] < x {
		this.maxStack = append(this.maxStack, x)
	} else {
		this.maxStack = append(this.maxStack, this.maxStack[len(this.maxStack)-1])
	}
	this.stack = append(this.stack, x)
}

func (this *MaxStack) Pop() int {
	out := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.maxStack = this.maxStack[:len(this.maxStack)-1]
	return out
}

func (this *MaxStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MaxStack) PeekMax() int {
	return this.maxStack[len(this.maxStack)-1]
}

func (this *MaxStack) PopMax() int {
	buffer := []int{}
	max := this.PeekMax()
	for max != this.Top() {
		buffer = append(buffer, this.Pop())
	}
	this.Pop()
	for i := len(buffer) - 1; i >= 0; i-- {
		this.Push(buffer[i])
	}
	return max
}
