package construct_smallest_number_from_di_string

import "testing"

func Test_smallestNumber(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"D"}, "21"},
		{"", args{"I"}, "12"},
		{"", args{"DD"}, "321"},
		{"", args{"II"}, "123"},
		{"", args{"DIDI"}, "21435"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestNumber(tt.args.pattern); got != tt.want {
				t.Errorf("smallestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
