package leetcode

import (
	"fmt"

	"github.com/hoshiin/leetcode-go/structures"
)

func FindDuplicateSubtrees(root *structures.TreeNode) []*structures.TreeNode {
	hashAll := map[string]int{}
	duplicate := []*structures.TreeNode{}
	dfs(root, hashAll, &duplicate)
	return duplicate
}

func dfs(node *structures.TreeNode, hashAll map[string]int, duplicate *[]*structures.TreeNode) string {
	if node == nil {
		return "nil"
	}
	lString := dfs(node.Left, hashAll, duplicate)
	rString := dfs(node.Right, hashAll, duplicate)
	buildString := fmt.Sprintf("[%s][%v][%s]", lString, node.Val, rString)
	hashAll[buildString]++
	if hashAll[buildString] == 2 {
		*duplicate = append(*duplicate, node)
	}
	return buildString
}
