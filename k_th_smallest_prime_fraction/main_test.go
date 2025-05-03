package k_th_smallest_prime_fraction

import (
	"reflect"
	"testing"
)

func Test_kthSmallestPrimeFraction(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k    int
		want []int
	}{
		{"", []int{1, 2, 3, 5}, 3, []int{2, 5}},
		{"", []int{1, 7}, 1, []int{1, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallestPrimeFraction(tt.arr, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kthSmallestPrimeFraction() = %v, want %v", got, tt.want)
			}
		})
	}
}
