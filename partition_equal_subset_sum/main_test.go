package partition_equal_subset_sum

import "testing"

func Test_canPartition(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"test1", []int{1, 5, 11, 5}, true},
		{"test2", []int{1, 2, 3, 5}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.nums); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
