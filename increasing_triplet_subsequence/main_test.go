package increasing_triplet_subsequence

import "testing"

func Test_increasingTriplet(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{[]int{1, 2, 3, 4, 5}}, true},
		{"", args{[]int{5, 4, 3, 2, 1}}, false},
		{"", args{[]int{2, 1, 5, 0, 4, 6}}, true},
		{"", args{[]int{5, 1, 5, 5, 2, 5, 4}}, true},
		{"", args{[]int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingTriplet(tt.args.nums); got != tt.want {
				t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}
