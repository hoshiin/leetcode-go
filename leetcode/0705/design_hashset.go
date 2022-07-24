package leetcode

import "github.com/hoshiin/leetcode-go/structures"

type MyHashSet struct {
	m [1024]*structures.ListNode
}

func Constructor() MyHashSet {
	return MyHashSet{[1024]*structures.ListNode{}}
}

func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}
	head := this.m[key%1024]
	if head == nil {
		this.m[key%1024] = &structures.ListNode{Val: key, Next: nil}
	} else {
		this.m[key%1024] = &structures.ListNode{Val: key, Next: head}
	}
}

func (this *MyHashSet) Remove(key int) {
	head := this.m[key%1024]
	if head == nil {
		return
	}
	// if key is at the front
	if head.Val == key {
		this.m[key%1024] = head.Next
		return
	}

	var prev *structures.ListNode
	for head != nil {
		if head.Val == key {
			prev.Next = head.Next
		}
		prev = head
		head = head.Next
	}
}

func (this *MyHashSet) Contains(key int) bool {
	head := this.m[key%1024]
	for head != nil {
		if head.Val == key {
			return true
		}
		head = head.Next
	}
	return false
}
