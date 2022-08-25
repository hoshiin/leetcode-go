package leetcode

import (
	"math"

	"github.com/hoshiin/leetcode-go/structures"
)

func ClosestValue(root *structures.TreeNode, target float64) int {
	closest := root.Val
	for root != nil {
		val := root.Val
		if math.Abs(float64(val)-target) < math.Abs(float64(closest)-target) {
			closest = val
		}
		if target < float64(val) {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return closest
}
