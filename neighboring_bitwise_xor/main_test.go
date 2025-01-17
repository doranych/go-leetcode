package neighboring_bitwise_xor

import "testing"

func Test_doesValidArrayExist(t *testing.T) {
	type args struct {
		derived []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{[]int{1, 1, 0}}, true},
		{"", args{[]int{1, 1}}, true},
		{"", args{[]int{1, 0}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doesValidArrayExist(tt.args.derived); got != tt.want {
				t.Errorf("doesValidArrayExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
