package make_lexicographically_smallest_array_by_swapping_elements

import (
	"reflect"
	"testing"
)

func Test_lexicographicallySmallestArray(t *testing.T) {
	type args struct {
		nums  []int
		limit int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{1, 5, 3, 9, 8}, 2}, []int{1, 3, 5, 8, 9}},
		{"", args{[]int{1, 7, 6, 18, 2, 1}, 3}, []int{1, 6, 7, 18, 1, 2}},
		{"", args{[]int{1, 7, 28, 19, 10}, 3}, []int{1, 7, 28, 19, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lexicographicallySmallestArray(tt.args.nums, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexicographicallySmallestArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
