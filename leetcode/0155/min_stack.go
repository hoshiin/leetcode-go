package leetcode

type MinStack struct {
	min   *int
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		min:   nil,
		stack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	if this.min == nil || *this.min > val {
		this.min = &val
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if this.stack[len(this.stack)-1] == *this.min {
		var minStack *int
		for _, n := range this.stack[:len(this.stack)-1] {
			if minStack == nil || *minStack > n {
				val := n
				minStack = &val
			}
		}
		this.min = minStack
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return *this.min
}
