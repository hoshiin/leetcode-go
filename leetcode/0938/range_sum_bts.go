package leetcode

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RangeSumBST(root *TreeNode, low int, high int) int {
	ans := 0
	dfs(root, low, high, &ans)
	return ans
}

func dfs(node *TreeNode, low int, high int, ans *int) {
	if node != nil {
		if low <= node.Val && node.Val <= high {
			*ans += node.Val
		}
		if low < node.Val {
			dfs(node.Left, low, high, ans)
		}
		if node.Val < high {
			dfs(node.Right, low, high, ans)
		}
	}
}

func RangeSumBSTIterative(root *TreeNode, low int, high int) int {
	ans := 0
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			if low <= node.Val && node.Val <= high {
				ans += node.Val
			}
			if low < node.Val {
				stack = append(stack, node.Left)
			}
			if node.Val < high {
				stack = append(stack, node.Right)
			}
		}

	}
	return ans
}
