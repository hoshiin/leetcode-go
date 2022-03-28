package leetcode

import "math"

type MovingAverage struct {
	windowSum int
	size      int
	head      int
	count     int
	queue     []int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size:  size,
		queue: make([]int, size),
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.count++
	tail := (this.head + 1) % this.size
	this.windowSum = this.windowSum - this.queue[tail] + val
	this.head = (this.head + 1) % this.size
	this.queue[this.head] = val
	return float64(this.windowSum) / math.Min(float64(this.size), float64(this.count))
}
