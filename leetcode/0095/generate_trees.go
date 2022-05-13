package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func GenerateTrees(n int) []*structures.TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(min, max int) []*structures.TreeNode {
	res := []*structures.TreeNode{}
	for i := min; i <= max; i++ {
		left := helper(min, i-1)
		right := helper(i+1, max)
		for _, l := range left {
			for _, r := range right {
				res = append(res, &structures.TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	if len(res) == 0 {
		res = append(res, nil)
	}
	return res
}
