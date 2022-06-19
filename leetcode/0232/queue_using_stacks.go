package leetcode

type MyQueue struct {
	inorder []int
	reverse []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	for len(this.reverse) != 0 {
		this.inorder = append(this.inorder, this.reverse[len(this.reverse)-1])
		this.reverse = this.reverse[:len(this.reverse)-1]
	}
	this.inorder = append(this.inorder, x)
}

func (this *MyQueue) Pop() int {
	for len(this.inorder) != 0 {
		this.reverse = append(this.reverse, this.inorder[len(this.inorder)-1])
		this.inorder = this.inorder[:len(this.inorder)-1]
	}
	pop := this.reverse[len(this.reverse)-1]
	this.reverse = this.reverse[:len(this.reverse)-1]
	return pop
}

func (this *MyQueue) Peek() int {
	if len(this.reverse) == 0 {
		return this.inorder[0]
	}
	return this.reverse[len(this.reverse)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.inorder) == 0 && len(this.reverse) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
