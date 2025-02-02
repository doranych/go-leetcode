package check_if_array_is_sorted_and_rotated

import (
	"fmt"
	"testing"
)

func Test_check(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{[]int{3, 4, 5, 1, 2}}, true},
		{"", args{[]int{2, 1, 3, 4}}, false},
		{"", args{[]int{1, 2, 3}}, true},
		{"", args{[]int{6, 7, 2, 7, 5}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.nums); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotation(t *testing.T) {
	A := []int{1, 2, 3, 4}
	B := make([]int, len(A))
	x := 2
	for i := 0; i < len(A); i++ {
		B[(i+x)%len(A)] = A[i]
	}
	fmt.Println(B)
}
