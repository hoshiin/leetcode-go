package leetcode

import "math/rand"

type RandomizedSet struct {
	set map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{set: map[int]int{}, arr: []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	}
	this.set[val] = len(this.arr)
	this.arr = append(this.arr, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	i, ok := this.set[val]
	if !ok {
		return false
	}
	this.set[this.arr[len(this.arr)-1]] = i
	delete(this.set, val)
	this.arr[i] = this.arr[len(this.arr)-1]
	this.arr = this.arr[:len(this.arr)-1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
