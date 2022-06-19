package leetcode

type MyStack struct {
	q1 []int
	q2 []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.q2 = append(this.q2, x)
	for len(this.q1) != 0 {
		this.q2 = append(this.q2, this.q1[0])
		this.q1 = this.q1[1:]
	}
	tmp := this.q1
	this.q1 = this.q2
	this.q2 = tmp
}

func (this *MyStack) Pop() int {
	pop := this.q1[0]
	this.q1 = this.q1[1:]
	return pop
}

func (this *MyStack) Top() int {
	return this.q1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.q1) == 0 && len(this.q2) == 0
}
