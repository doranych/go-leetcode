package largest_divisible_subset

import (
	"reflect"
	"testing"
)

func Test_largestDivisibleSubset(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"", []int{1, 2, 3}, []int{1, 3}},
		{"", []int{1, 2, 4, 8}, []int{1, 2, 4, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestDivisibleSubset(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestDivisibleSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
