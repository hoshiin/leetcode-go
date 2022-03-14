package leetcode

import (
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		head *structures.ListNode
		want bool
	}{
		{structures.NewListNode([]int{1, 2, 2, 1}), true},
		{structures.NewListNode([]int{1, 2}), false},
	}
	for _, tt := range tests {
		if got := IsPalindrome(tt.head); got != tt.want {
			t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
		}
	}
}
