package leetcode

import "fmt"

type TwoSum struct {
	nums map[int]int
}

func Constructor() TwoSum {
	return TwoSum{nums: map[int]int{}}
}

func (this *TwoSum) Add(number int) {
	this.nums[number]++
}

func (this *TwoSum) Find(value int) bool {
	for num, count := range this.nums {
		if _, ok := this.nums[value-num]; ok && (value-num != num || count > 1) {
			return true
		}
	}
	return false
}

func TestTwoSum() {
	obj := Constructor()
	obj.Add(1)               // [] --> [1]
	obj.Add(3)               // [1] --> [1,3]
	obj.Add(5)               // [1,3] --> [1,3,5]
	fmt.Println(obj.Find(4)) // true
	fmt.Println(obj.Find(7)) // false
}
