package find_the_prefix_common_array_of_two_arrays

import (
	"reflect"
	"testing"
)

func Test_findThePrefixCommonArray(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{1, 3, 2, 4}, []int{3, 1, 2, 4}}, []int{0, 2, 3, 4}},
		{"", args{[]int{2, 3, 1}, []int{3, 1, 2}}, []int{0, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findThePrefixCommonArray(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findThePrefixCommonArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
