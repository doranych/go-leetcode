package merge_two_2d_arrays_by_summing_values

import (
	"reflect"
	"testing"
)

func Test_mergeArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 [][]int
		nums2 [][]int
		want  [][]int
	}{
		{
			"",
			[][]int{{1, 3}, {2, 4}, {3, 5}},
			[][]int{{1, 2}, {2, 3}, {3, 4}},
			[][]int{{1, 5}, {2, 7}, {3, 9}},
		},
		{
			"",
			[][]int{{1, 3}, {2, 4}, {3, 5}},
			[][]int{{4, 2}, {5, 3}, {6, 4}},
			[][]int{{1, 3}, {2, 4}, {3, 5}, {4, 2}, {5, 3}, {6, 4}},
		},
		{
			"",
			[][]int{{2, 4}, {3, 6}, {5, 5}},
			[][]int{{1, 3}, {4, 3}},
			[][]int{{1, 3}, {2, 4}, {3, 6}, {4, 3}, {5, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeArrays(tt.nums1, tt.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
