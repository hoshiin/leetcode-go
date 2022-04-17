package leetcode

import (
	"reflect"
	"testing"
)

func TestCloneGraph(t *testing.T) {
	n1 := Node{Val: 1}
	n2 := Node{Val: 2}
	n3 := Node{Val: 3}
	n4 := Node{Val: 4}
	n1.Neighbors = []*Node{&n2, &n4}
	n2.Neighbors = []*Node{&n1, &n3}
	n3.Neighbors = []*Node{&n2, &n4}
	n4.Neighbors = []*Node{&n1, &n3}
	tests := []struct {
		node *Node
		want *Node
	}{
		{&n1, &n1},
		{&Node{}, &Node{}},
		{&Node{Val: 1}, &Node{Val: 1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := CloneGraph(tt.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CloneGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
