package leetcode

import (
	"math"

	"github.com/hoshiin/leetcode-go/structures"
)

func MaxDepth(root *structures.TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1

}

func MaxDepthTopDown(root *structures.TreeNode) int {
	answer := 0
	maxDepthInternal(root, 0, &answer)
	return answer
}

func maxDepthInternal(root *structures.TreeNode, depth int, answer *int) {
	if root == nil {
		*answer = int(math.Max(float64(depth), float64(*answer)))
		return
	}
	maxDepthInternal(root.Left, depth+1, answer)
	maxDepthInternal(root.Right, depth+1, answer)
}
