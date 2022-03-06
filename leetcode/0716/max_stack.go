package leetcode

import "math"

type MaxStack struct {
	stack     []int
	lastIndex int
	maxIndex  int
}

func Constructor() MaxStack {
	return MaxStack{
		stack:     []int{},
		lastIndex: -1,
		maxIndex:  -1,
	}
}

func (this *MaxStack) Push(x int) {
	if len(this.stack) == 0 || x >= this.stack[this.maxIndex] {
		this.maxIndex = len(this.stack)
	}
	this.stack = append(this.stack, x)
	this.lastIndex++
}

func (this *MaxStack) Pop() int {
	if this.lastIndex == this.maxIndex {
		return this.PopMax()
	}
	out := this.stack[this.lastIndex]
	this.stack = this.stack[:this.lastIndex]
	this.lastIndex--
	return out
}

func (this *MaxStack) Top() int {
	return this.stack[this.lastIndex]
}

func (this *MaxStack) PeekMax() int {
	return this.stack[this.maxIndex]
}

func (this *MaxStack) PopMax() int {
	out := this.stack[this.maxIndex]
	if this.maxIndex+1 <= len(this.stack)-1 {
		this.stack = append(this.stack[:this.maxIndex], this.stack[this.maxIndex+1:]...)
	} else {
		this.stack = this.stack[:this.maxIndex]
	}
	this.lastIndex--
	max := int(math.Inf(-1))
	for i, s := range this.stack {
		if s >= max {
			max = s
			this.maxIndex = i
		}
	}

	return out
}
