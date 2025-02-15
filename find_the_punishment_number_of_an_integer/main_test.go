package find_the_punishment_number_of_an_integer

import "testing"

func Test_punishmentNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{10}, 182},
		{"", args{37}, 1478},
		{"", args{123}, 41334},
		{"", args{321}, 184768},
		{"", args{657}, 1204515},
		{"", args{1000}, 10804657},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := punishmentNumber(tt.args.n); got != tt.want {
				t.Errorf("punishmentNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
