package battleships_in_a_board

import "testing"

func Test_countBattleships(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test case 1", args{[][]byte{
			{'X', '.', '.', 'X'},
			{'.', '.', '.', 'X'},
			{'.', '.', '.', 'X'},
		}}, 2},
		{"Test case 2", args{[][]byte{
			{'.'},
		}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBattleships(tt.args.board); got != tt.want {
				t.Errorf("countBattleships() = %v, want %v", got, tt.want)
			}
		})
	}
}
