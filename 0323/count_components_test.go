package leetcode

import "testing"

func TestCountComponents(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{5, [][]int{{0, 1}, {1, 2}, {3, 4}}}, 2},
		{"", args{5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountComponents(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("CountComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}
