package maximum_employees_to_be_invited_to_a_meeting

import "testing"

func Test_maximumInvitations(t *testing.T) {
	type args struct {
		favorite []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{2, 2, 1, 2}}, 3},
		{"", args{[]int{1, 2, 0}}, 3},
		{"", args{[]int{3, 0, 1, 4, 1}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumInvitations(tt.args.favorite); got != tt.want {
				t.Errorf("maximumInvitations() = %v, want %v", got, tt.want)
			}
		})
	}
}
