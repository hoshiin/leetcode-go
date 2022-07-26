package leetcode

import (
	"reflect"
	"testing"
)

func TestFindRestaurant(t *testing.T) {
	tests := []struct {
		list1 []string
		list2 []string
		want  []string
	}{
		{[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}, []string{"Shogun"}},
		{[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}, []string{"Shogun"}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FindRestaurant(tt.list1, tt.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindRestaurant() = %v, want %v", got, tt.want)
			}
		})
	}
}
