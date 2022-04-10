package leetcode

type MyCircularQueue struct {
	queue []int
	head  int
	tail  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k),
		head:  -1,
		tail:  -1,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.IsEmpty() {
		this.head = 0
	}
	this.tail = (this.tail + 1) % len(this.queue)
	this.queue[this.tail] = value
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.head == this.tail {
		this.head, this.tail = -1, -1
		return true
	}
	this.head = (this.head + 1) % len(this.queue)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == -1
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.tail+1)%len(this.queue) == this.head
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
