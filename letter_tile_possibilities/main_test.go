package letter_tile_possibilities

import "testing"

func Test_numTilePossibilities(t *testing.T) {
	type args struct {
		tiles string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"A"}, 1},
		{"", args{"AA"}, 2},
		{"", args{"AAB"}, 8},
		{"", args{"AAABBC"}, 188},
		{"", args{"V"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTilePossibilities(tt.args.tiles); got != tt.want {
				t.Errorf("numTilePossibilities() = %v, want %v", got, tt.want)
			}
		})
	}
}
