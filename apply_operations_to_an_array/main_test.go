package apply_operations_to_an_array

import (
	"reflect"
	"testing"
)

func Test_applyOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"", []int{1, 2, 2, 1, 1, 0}, []int{1, 4, 2, 0, 0, 0}},
		{"", []int{0, 1}, []int{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyOperations(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
