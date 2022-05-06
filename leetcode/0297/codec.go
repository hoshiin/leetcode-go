package leetcode

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hoshiin/leetcode-go/structures"
)

type Codec struct {
	input []string
}

func Constructor() Codec {
	return Codec{input: []string{}}
}

// Serializes a tree to a single string.
func (this *Codec) Serialize(root *structures.TreeNode) string {
	return this.serializeHelper(root, "")
}

func (this *Codec) serializeHelper(root *structures.TreeNode, str string) string {
	if root == nil {
		str = fmt.Sprintf("%snull,", str)
	} else {
		str = fmt.Sprintf("%s%s,", str, strconv.Itoa(root.Val))
		str = this.serializeHelper(root.Left, str)
		str = this.serializeHelper(root.Right, str)
	}
	return str
}

// Deserializes your encoded data to tree.
func (this *Codec) Deserialize(data string) *structures.TreeNode {
	this.input = strings.Split(data, ",")
	return this.deserializeHelper()
}

func (this *Codec) deserializeHelper() *structures.TreeNode {
	if this.input[0] == "null" {
		this.input = this.input[1:]
		return nil
	}
	val, _ := strconv.Atoi(this.input[0])
	this.input = this.input[1:]

	root := &structures.TreeNode{Val: val}
	root.Left = this.deserializeHelper()
	root.Right = this.deserializeHelper()
	return root
}
